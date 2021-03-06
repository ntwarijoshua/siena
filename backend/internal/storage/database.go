package storage

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ntwarijoshua/siena/internal/models"
	logger "github.com/sirupsen/logrus"
)

// DBCredentials credentials structure
type DBCredentials struct {
	Host     string
	Port     string
	Dbname   string
	Username string
	Password string
}

func NewDB(config DBCredentials) (*models.DataStore, error) {
	logger.Info("Connecting to the database")
	database, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		config.Host, config.Port, config.Dbname, config.Username, config.Password))
	return &models.DataStore{DB: database}, err
}
