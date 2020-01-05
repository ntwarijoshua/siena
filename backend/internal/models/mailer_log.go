package models

import (
	logger "github.com/sirupsen/logrus"
	"time"
)

var SupportedStatus = map[string]string{
	"PROCESSING": "processing",
	"QUEUED":     "queued",
	"SENT":       "sent",
}

var SupportedMessageType = map[string]string{
	"CONFIRMATION":   "confirmation_mail",
	"PASSWORD_RESET": "password_reset_mail",
}

type MailerLog struct {
	Id        int       `db:"id"`
	Type      string    `db:"type"`
	Payload   string    `db:"payload"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (dataStore *DataStore) SaveMailLog(log MailerLog) (MailerLog, error) {
	var (
		createMailLogQuery = `
			INSERT INTO mailer_logs (type, payload, status, created_at) VALUES
			(:type, :payload, :status, :created_at) RETURNING *;`
	)
	rows, err := dataStore.DB.NamedQuery(createMailLogQuery, &log)
	if err != nil {
		logger.Errorf("Error occurred while try to save a mail log", err)
	}
	for rows.Next() {
		if err = rows.StructScan(&log); err != nil {
			logger.Errorf("Error occurred while try to save a mail log", err)
		}
	}
	return log, err
}

func (dataStore *DataStore) UpdateMailLog(log MailerLog) (MailerLog, error) {
	var (
		updateMailLogQuery = `UPDATE mailer_logs SET status = :status WHERE id = :id RETURNING *;`
	)
	rows, err := dataStore.DB.NamedQuery(updateMailLogQuery, &log)
	if err != nil {
		logger.Errorf("Error occurred while try to update a mail log", err)
	}
	for rows.Next() {
		if err = rows.StructScan(&log); err != nil {
			logger.Errorf("Error occurred while try to update a mail log", err)
		}
	}
	return log, err
}
