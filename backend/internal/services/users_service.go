package services

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nsqio/go-nsq"
	"github.com/ntwarijoshua/siena/internal/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

const ClientRoleSlug = "client"
const ConfirmationMailTopic = "account-confirmation-emails"
const ConfirmationMailChannel = "account-confirmation-channel"

type UserTransactionMessage struct {
	Name string					`json:"name"`
	EmailAddress string			`json:"email_address"`
	Token string				`json:"token"`
	Subject string				`json:"subject"`
	TrackingId int				`json:"tracking_id"`
	Type string					`json:"type"`
}

type CustomClaims struct {
	UserId int `json:"userId"`
	Email string	`json:"email"`
	IssuedAt time.Time
	jwt.StandardClaims
}

type UserService struct{
	dataLayer *models.DataStore
	logger *logrus.Logger
}

func (s *UserService)CreateUser(user models.User, profile models.Profile) (models.User, error) {
	role, err  := s.dataLayer.GetRoleBySlug(ClientRoleSlug)
	if err != nil {
		s.logger.Errorf("Error while creating a user:", err)
		return user, err
	}
	hashAndSalt, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		s.logger.Errorf("Error while hashing user password:", err)
		return user, err
	}
	user.Password = string(hashAndSalt)
	user, err = s.dataLayer.SaveUser(profile, user, role)
	if err != nil {
		return user,err
	}
	// queue confirmation mail
	confirmationToken, err := generateUniqueTokenForUser(user)
	if err != nil {
		s.logger.Errorf("Error occurred while generating a confirmation token", err)
		return user,err
	}
	confirmationMessage := UserTransactionMessage{
		Name: profile.Names,
		EmailAddress:     	user.Email,
		Token: confirmationToken,
		Subject: "Confirm your account!",
	}
	rawPayload, _ := json.Marshal(confirmationMessage)
	messageLog := models.MailerLog{
		Type:      models.SupportedMessageType["CONFIRMATION"],
		Payload:   string(rawPayload),
		Status:    models.SupportedStatus["PROCESSING"],
		CreatedAt: time.Now(),
	}
	messageLog,err = s.dataLayer.SaveMailLog(messageLog)
	if err != nil {
		s.logger.Errorf("Error occurred while queueing user confirmation mail", err)
	}
	confirmationMessage.TrackingId = messageLog.Id
	confirmationMessage.Type = messageLog.Type
	if err = produceConfirmationMail(confirmationMessage); err != nil {
		s.logger.Errorf("Error occurred while queueing user confirmation mail", err)
		return user, err
	}
	messageLog.Status = models.SupportedStatus["QUEUED"]
	messageLog, err = s.dataLayer.UpdateMailLog(messageLog)

	return user, err
}

func (s *UserService) GetUserByMail(email string) (models.User, error) {
	return s.dataLayer.GetUserByEmail(email)
}
func (s *UserService) GetJWTToken(user models.User, password string) (string, error)  {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}
	signingKey := os.Getenv("JWT_SIGNING_KEY")

	claims := CustomClaims{
		user.Id,
		user.Email,
		time.Now(),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 45).Unix(),
			Issuer: "api.siena",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

func produceConfirmationMail (message UserTransactionMessage) error {
	config := nsq.NewConfig()
	w, err := nsq.NewProducer(os.Getenv("NSQD"), config)
	defer func() {_ = w.Stop}()
	if err != nil {
		return err
	}
	rawMSG, err := json.Marshal(message);
	if err != nil {
		return err
	}
	return w.Publish(ConfirmationMailTopic, []byte(rawMSG))
}

func generateUniqueTokenForUser(user models.User) (string,error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "",err
	}
	leading := fmt.Sprintf("account-%d", user.Id)
	token := fmt.Sprintf("%s-%x-%x-%x-%x-%x", leading, b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return token, nil
}

