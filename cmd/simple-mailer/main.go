package main

import (
	"flag"
	"log"

	"github.com/Aditya0812/simple-mailer/internal/config"
	"github.com/Aditya0812/simple-mailer/internal/mailer"
	"github.com/Aditya0812/simple-mailer/internal/parser"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading the .env file. ", err)
	}

	config.Init()

	excelFilePath := flag.String("d", "../../data/data.xlsx", "An Excel file containing the recipient details.")
	templateFilePath := flag.String("t", "../../templates/email_template.html", "An HTML file containing the Email template.")
	subject := flag.String("s", "", "Subject of the Email to be sent.")
	flag.Parse()

	recipients, err := parser.ParseExcel(*excelFilePath, "Sheet1")
	if err != nil {
		log.Fatalf("Failed to parse Excel file: %v", err)
	}

	if err = mailer.SendBulkEmails(*templateFilePath, recipients, *subject); err != nil {
		log.Fatalf("Error sending emails: %v", err)
	}

	log.Println("Emails sent successfully!")
}
