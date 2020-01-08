package services

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nsqio/go-nsq"
	"github.com/ntwarijoshua/siena/internal/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

const ClientRoleSlug = "client"
const ConfirmationMailTopic = "account-confirmation-emails"
const ConfirmationMailChannel = "account-confirmation-channel"

type UserTransactionMessage struct {
	Name         string `json:"name"`
	EmailAddress string `json:"email_address"`
	Token        string `json:"token"`
	Subject      string `json:"subject"`
	TrackingId   int    `json:"tracking_id"`
	Type         string `json:"type"`
}

type CustomClaims struct {
	UserId   int    `json:"userId"`
	Email    string `json:"email"`
	IssuedAt time.Time
	jwt.StandardClaims
}

type UserService struct {
	dataLayer *models.DataStore
	userModel *models.User
	roleModel *models.Role
	logger    *logrus.Logger
	context   context.Context
}

func (s *UserService) CreateUser(user models.User, profile models.Profile) (models.User, error) {
	var err error
	clientRole, err := models.Roles(qm.Where("slug = ?", ClientRoleSlug)).One(s.context, s.dataLayer.DB)
	if err != nil {
		s.logger.Errorf("Error occurred %s", errors.Cause(err))
	}
	if err = user.SetRole(s.context, s.dataLayer.DB, false, clientRole); err != nil {
		s.logger.Errorf("Error occurred %s", errors.Cause(err))
	}

	if err = user.SetProfile(s.context, s.dataLayer.DB, true, &profile); err != nil {
		s.logger.Errorf("Error occurred %s", errors.Cause(err))
	}

	hashAndSalt, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		s.logger.Errorf("Error while hashing user password:", err)
		return user, err
	}
	user.Password = string(hashAndSalt)
	err = user.Insert(s.context, s.dataLayer.DB, boil.Infer())
	if err != nil {
		s.logger.Errorf("Could not create user %s", errors.Cause(err))
		return user, err
	}

	// queue confirmation mail
	confirmationToken, err := generateUniqueTokenForUser(user)
	if err != nil {
		s.logger.Errorf("Error occurred while generating a confirmation token %s", errors.Cause(err))
		return user, err
	}
	confirmationMessage := UserTransactionMessage{
		Name:         user.R.Profile.Names.String,
		EmailAddress: user.Email,
		Token:        confirmationToken,
		Subject:      "Confirm your account!",
	}
	rawPayload, _ := json.Marshal(confirmationMessage)
	messageLog := models.MailerLog{
		Type:      models.SupportedMessageType["CONFIRMATION"],
		Payload:   string(rawPayload),
		Status:    models.SupportedStatus["PROCESSING"],
		CreatedAt: time.Now(),
	}
	err = messageLog.Insert(s.context, s.dataLayer.DB, boil.Infer())
	if err != nil {
		s.logger.Errorf("Error occurred while queueing user confirmation mail %s", errors.Cause(err))
	}
	confirmationMessage.TrackingId = messageLog.ID
	confirmationMessage.Type = messageLog.Type
	if err = produceConfirmationMail(confirmationMessage); err != nil {
		s.logger.Errorf("Error occurred while queueing user confirmation mail %s", errors.Cause(err))
		return user, err
	}
	messageLog.Status = models.SupportedStatus["QUEUED"]
	messageLog.Payload = string(func(message interface{}) []byte {
		marshalled, _ := json.Marshal(message)
		return marshalled
	}(confirmationMessage))
	_, err = messageLog.Update(s.context, s.dataLayer.DB, boil.Infer())
	if err != nil {
		s.logger.Errorf("Error occurred %s", errors.Cause(err))
		return user, err
	}
	return user, err
}

func (s *UserService) GetUserByMail(email string) (*models.User, error) {
	return models.Users(qm.Where("email = ?", email)).One(s.context, s.dataLayer.DB)
}
func (s *UserService) GetJWTToken(user *models.User, password string) (string, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}
	signingKey := os.Getenv("JWT_SIGNING_KEY")

	claims := CustomClaims{
		user.ID,
		user.Email,
		time.Now(),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 45).Unix(),
			Issuer:    "api.siena",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

func produceConfirmationMail(message UserTransactionMessage) error {
	config := nsq.NewConfig()
	w, err := nsq.NewProducer(os.Getenv("NSQD"), config)
	defer func() { _ = w.Stop }()
	if err != nil {
		return err
	}
	rawMSG, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return w.Publish(ConfirmationMailTopic, []byte(rawMSG))
}

func generateUniqueTokenForUser(user models.User) (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	leading := fmt.Sprintf("account-%d", user.ID)
	token := fmt.Sprintf("%s-%x-%x-%x-%x-%x", leading, b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return token, nil
}
