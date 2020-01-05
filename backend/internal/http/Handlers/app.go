package Handlers

import (
	"github.com/ntwarijoshua/siena/internal/services"
	"github.com/sirupsen/logrus"
)

type App struct {
	Logger *logrus.Logger
	ServiceContainer *services.ServiceContainer
}
