package models

// SupportedMessageType is the types of messages we can send out
var SupportedMessageType = map[string]string{
	"CONFIRMATION":   "confirmation_mail",
	"PASSWORD_RESET": "password_reset_mail",
}
// SupportedStatus message statuses a message can have
var SupportedStatus = map[string]string{
	"PROCESSING": "processing",
	"QUEUED":     "queued",
	"SENT":       "sent",
}
