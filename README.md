# SMTP-Mailer

# Simple Mailer

This is a simple bulk mail sender CLI application, built using Go.

## Setup

## Procedure to run the application

In the root directory, run the following command:

```bash
go run ./cmd/simple-mailer/main.go -d <.xlsx_file_path> -t <html_template_file_path> -s <subject_of_mail>
```

Example

```bash

go run ./cmd/simple-mailer/main.go -d "./data/data.xlsx" -t "./templates/email_template.html" -s "Hello World"
```
