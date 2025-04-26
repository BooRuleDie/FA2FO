package mailer

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendgridMailer struct {
	fromEmail string
	apiKey    string
	client    *sendgrid.Client
}

func NewSendgrid(apiKey, fromEmail string) *SendgridMailer {
	client := sendgrid.NewSendClient(apiKey)

	return &SendgridMailer{
		fromEmail: fromEmail,
		apiKey:    apiKey,
		client:    client,
	}
}

func (m *SendgridMailer) Send(templateFile, username, email string, data any, isSandbox bool) error {
	from := mail.NewEmail(FromName, m.fromEmail)
	to := mail.NewEmail(username, email)

	// template parsing and building
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return err
	}

	message := mail.NewSingleEmail(from, subject.String(), to, "", body.String())

	message.SetMailSettings(&mail.MailSettings{
		SandboxMode: &mail.Setting{
			Enable: &isSandbox,
		},
	})

	var retryErr error
	for i := range maxRetries {
		_, err := m.client.Send(message)
		if err != nil {
			if retryErr == nil {
				retryErr = fmt.Errorf("attempt %d: %w", i+1, err)
			} else {
				retryErr = fmt.Errorf("%v, attempt %d: %w", retryErr, i+1, err)
			}

			// exponential backoff
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}

		return nil
	}

	return fmt.Errorf("failed to send email after %d attempts. Errors: %v", maxRetries, retryErr)
}
