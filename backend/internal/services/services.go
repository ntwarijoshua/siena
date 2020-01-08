package services

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v3"
	"github.com/ntwarijoshua/siena/internal/models"
	"github.com/sirupsen/logrus"
	"os"
)

type ServiceContainer struct {
	services map[string]interface{}
	Store    *models.DataStore
	Logger   *logrus.Logger
	Context  context.Context
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
		"userService":       NewUserUserService(sc.Context, sc.Store, sc.Logger),
		"mailerService":     NewMailerService(sc.Context, sc.Store, sc.Logger),
		"validationService": NewValidationService(sc.Context, sc.Store, sc.Logger),
	}
}

func NewUserUserService(context context.Context, store *models.DataStore, logger *logrus.Logger) *UserService {
	return &UserService{
		dataLayer: store,
		userModel: &models.User{},
		roleModel: &models.Role{},
		logger:    logger,
		context:   context,
	}
}

func NewMailerService(context context.Context, store *models.DataStore, logger *logrus.Logger) *MailerService {
	// resolving service dependencies
	mailGunClientInstance := mailgun.NewMailgun(
		os.Getenv("MAIL_GUN_DOMAIN"), os.Getenv("MAIL_GUN_API_KEY"),
	)
	return &MailerService{
		MGClient: mailGunClientInstance,
		logger:   logger,
		store:    store,
		context:  context,
	}
}

func NewValidationService(context context.Context, store *models.DataStore, logger *logrus.Logger) *ValidationService {
	validationService := ValidationService{
		dataLayer: store,
		logger:    logger,
		context:   context,
	}
	validationService.InitializeValidator()
	return &validationService
}
