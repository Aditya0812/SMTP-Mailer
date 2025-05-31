package mailer

import (
	"github.com/Aditya0812/simple-mailer/internal/config"
	"gopkg.in/gomail.v2"
)

func createMessage(receiverEmailID, subject, content string) *gomail.Message {
	m := gomail.NewMessage()

	m.SetHeaders(map[string][]string{
		"From":    {config.Conf.SenderEmailID},
		"To":      {receiverEmailID},
		"Subject": {subject},
	})

	m.SetBody("text/html", content)

	return m
}
