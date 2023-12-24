package utils

import (
	"github.com/ritesh-15/notesync-backend/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	to      string
	message string
	html    string
	subject string
}

func NewEmail(to string, message string, html string, subject string) *Email {
	return &Email{
		to:      to,
		message: message,
		html:    html,
		subject: subject,
	}
}

func SendEmail(email *Email) (int, error) {
	from := mail.NewEmail("notesync", config.SENDGRID_EMAIL)

	to := mail.NewEmail("", email.to)

	message := mail.NewSingleEmail(from, email.subject, to, email.message, email.html)

	client := sendgrid.NewSendClient(config.SENDGRID_API_KEY)

	response, err := client.Send(message)

	return response.StatusCode, err
}
