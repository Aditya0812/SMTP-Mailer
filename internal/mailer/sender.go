package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"sync"

	"github.com/Aditya0812/simple-mailer/internal/config"
	"github.com/Aditya0812/simple-mailer/internal/parser"
	"gopkg.in/gomail.v2"
)

func sendEmail(d *gomail.Dialer, tmpl *template.Template, recipient parser.Recipient, subject string) {
	var body bytes.Buffer

	if err := tmpl.Execute(&body, struct{ Name string }{Name: recipient.Name}); err != nil {
		log.Printf("Error processing template for %s: %v", recipient.Email, err)
		return
	}

	msg := createMessage(recipient.Email, subject, body.String())

	if err := d.DialAndSend(msg); err != nil {
		log.Printf("Error sending email to %s: %v", recipient.Email, err)
		return
	}

	log.Printf("Email send to %s", recipient.Email)
}

func SendBulkEmails(templateFilePath string, recipients []parser.Recipient, subject string) error {
	tmpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("error loading email template: %w", err)
	}

	d := newDialer()

	var wg sync.WaitGroup
	emailChan := make(chan parser.Recipient, len(recipients))

	for range config.Conf.WorkerCount {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for recipient := range emailChan {
				sendEmail(d, tmpl, recipient, subject)
			}
		}()
	}

	for _, recipient := range recipients {
		emailChan <- recipient
	}

	close(emailChan)
	wg.Wait()

	return nil
}
