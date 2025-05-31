package mailer

import (
	"github.com/Aditya0812/simple-mailer/internal/config"
	"gopkg.in/gomail.v2"
)

func newDialer() *gomail.Dialer {
	return gomail.NewDialer(config.Conf.SMTPHost, config.Conf.SMTPPort, config.Conf.SenderEmailID, config.Conf.SenderEmailPassword)
}
