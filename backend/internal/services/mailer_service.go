package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mailgun/mailgun-go/v3"
	"github.com/nsqio/go-nsq"
	"github.com/ntwarijoshua/siena/internal/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"os"
	"os/signal"
	"syscall"
)

type MailerService struct {
	MGClient *mailgun.MailgunImpl
	Message  *mailgun.Message
	logger   *logrus.Logger
	store    *models.DataStore
	context  context.Context
}

func (ms *MailerService) HandleMessage(m *nsq.Message) error {
	var message = UserTransactionMessage{}
	if err := json.Unmarshal(m.Body, &message); err != nil {
		ms.logger.Errorf("Failed to read message from nsq", err)
		return err
	}
	receiverEmail := message.EmailAddress
	senderEmail := os.Getenv("MAIL_SENDER")
	subject := message.Subject
	ms.Message = ms.MGClient.NewMessage(senderEmail, subject, "", receiverEmail)

	if message.Type == models.SupportedMessageType["CONFIRMATION"] {
		return ms.sendAccountConfirmationEmail(message)
	}
	return nil
}

func (ms *MailerService) sendAccountConfirmationEmail(msg UserTransactionMessage) error {
	ms.Message.SetTemplate("account-confirmation-email")
	_ = ms.Message.AddVariable(
		"confirmation_link",
		fmt.Sprintf("https://foobar.com?id=%d&token=%s", msg.TrackingId, msg.Token))

	_, _, err := ms.MGClient.Send(ms.context, ms.Message)
	if err != nil {
		ms.logger.Errorf("Error occured while trying to send out mail", err)
		return err
	}

	// Update the mail log as sent
	persistedMailLog, err := models.MailerLogs(qm.Where("id = ?", msg.TrackingId)).
		One(ms.context, ms.store.DB)
	if err != nil {
		ms.logger.Errorf("Unexpected error occurred", errors.Cause(err))
		return err
	}
	persistedMailLog.Status = models.SupportedStatus["SENT"]
	_, err = persistedMailLog.Update(ms.context, ms.store.DB, boil.Infer())
	if err != nil {
		ms.logger.Errorf("Unexpected error occurred", errors.Cause(err))
		return err
	}
	return nil
}

func StartMailerConsumer(logger *logrus.Logger, messagesHandler *MailerService) error {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(ConfirmationMailTopic, ConfirmationMailChannel, config)
	if err != nil {
		logger.Errorf("Error while trying to initialize consumer", err)
		return err
	}
	consumer.ChangeMaxInFlight(200)
	consumer.AddConcurrentHandlers(
		messagesHandler,
		20,
	)
	if err = consumer.ConnectToNSQLookupd(os.Getenv("NSQLOOKUPD")); err != nil {
		logger.Errorf("Error occurred while trying to connect to NSQD", err)
		return err
	}
	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT)

	for {
		select {
		case <-consumer.StopChan:
			return nil
		case <-shutdown:
			consumer.Stop()
		}
	}
}
