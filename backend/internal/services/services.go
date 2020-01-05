package services

import (
	"fmt"
	"github.com/mailgun/mailgun-go/v3"
	"github.com/ntwarijoshua/siena/internal/models"
	"github.com/sirupsen/logrus"
	"os"
)

type ServiceContainer struct {
	services  map[string]interface{}
	DataLayer *models.DataStore
	Logger *logrus.Logger
}

func (sc *ServiceContainer) GetService(serviceName string) interface{} {
	service, isRegistered := sc.services[serviceName]
	if !isRegistered {
		panic(fmt.Sprintf("Service %s is not registered", serviceName))
	}
	return service
}

func (sc *ServiceContainer) BuildServiceContainer() {
	sc.services = map[string]interface{}{
		"userService":   NewUserUserService(sc.DataLayer, sc.Logger),
		"mailerService": NewMailerService(sc.Logger),
		"validationService": NewValidationService(sc.DataLayer, sc.Logger),
	}
}

func NewUserUserService(layer *models.DataStore, logger *logrus.Logger) *UserService {
	return &UserService{dataLayer: layer, logger: logger}
}

func NewMailerService(logger *logrus.Logger) *MailerService {
	// resolving service dependencies
	mailGunClientInstance := mailgun.NewMailgun(
		os.Getenv("MAIL_GUN_DOMAIN"),os.Getenv("MAIL_GUN_API_KEY"),
		)
	return &MailerService{
		MGClient: mailGunClientInstance,
		logger: logger,
	}
}

func NewValidationService(layer *models.DataStore, logger *logrus.Logger) *ValidationService {
	validationService := ValidationService{dataLayer:layer, logger:logger}
	validationService.InitializeValidator()
	return &validationService
}
