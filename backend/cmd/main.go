package main

import (
	dotenv "github.com/joho/godotenv"
	internalHttp "github.com/ntwarijoshua/siena/internal/http"
	"github.com/ntwarijoshua/siena/internal/http/Handlers"
	"github.com/ntwarijoshua/siena/internal/models"
	"github.com/ntwarijoshua/siena/internal/services"
	"github.com/ntwarijoshua/siena/internal/storage"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)

func setupLogger() *logrus.Logger {
	logger := logrus.New()
	cwd, err := os.Getwd()
	if err != nil {
		logger.Fatalf("Failed to determine working directory: %s", err)
	}
	runID := time.Now().Format("run-2006-01-02")
	logDirPath := filepath.Join(cwd, "logs")
	if _, err := os.Stat(logDirPath); os.IsNotExist(err) {
		if err = os.Mkdir(logDirPath, 0777); err != nil {
			logger.Fatalf("Failed to create logs directory %s", err)
		}
	}
	logLocation := filepath.Join(logDirPath, runID+".log")
	logFile, err := os.OpenFile(logLocation, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		logger.Fatalf("Failed to create log file: %s", err)
	}
	logger.SetOutput(io.MultiWriter(os.Stderr, logFile))
	logrus.RegisterExitHandler(func() {
		if logFile == nil {
			return
		}
		if err = logFile.Close(); err != nil {
			logger.Fatalf("Failed to close logfile %s", err)
		}
	})
	return logger
}

func loadEnvVars() {
	err := dotenv.Load()
	if err != nil {
		logrus.Fatalf("Could not load environment variables %s", err)
	}
}



func initDB() *models.DataStore {
	dataLayer, err := storage.NewDB(storage.DBCredentials{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Dbname:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	})
	if err != nil {
		logrus.Fatalf("Could not initialize connection to database %s", err)
	}
	return dataLayer
}



func main() {
	loadEnvVars()
	appLogger := setupLogger()
	serviceContainer := &services.ServiceContainer{DataLayer: initDB(), Logger: appLogger}
	serviceContainer.BuildServiceContainer()
	app := Handlers.App{
		Logger: appLogger,
		ServiceContainer: serviceContainer,
	}
	go services.StartMailerConsumer(appLogger)
	defer app.Logger.Exit(0)
	app.Logger.Fatal(internalHttp.GetRouter(app).Run(":8090"))
}
