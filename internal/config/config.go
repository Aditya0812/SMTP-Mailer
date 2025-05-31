package config

import (
	"github.com/Aditya0812/simple-mailer/internal/env"
)

type conf struct {
	SMTPHost            string
	SMTPPort            int
	SenderEmailID       string
	SenderEmailPassword string
	WorkerCount         int
}

var Conf *conf

func Init() {
	Conf = &conf{
		SMTPHost:            smtpHost,
		SMTPPort:            smtpPort,
		SenderEmailID:       env.GetString("SENDER_EMAIL_ID", ""),
		SenderEmailPassword: env.GetString("SENDER_EMAIL_PASSWORD", ""),
		WorkerCount:         workerCount,
	}
}
