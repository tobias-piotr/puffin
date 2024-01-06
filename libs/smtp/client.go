package smtp

import (
	"bytes"
	"html/template"
	"os"

	"github.com/go-mail/mail"
)

type Dialer interface {
	DialAndSend(m ...*mail.Message) error
}

// SmtpClient represents a client for sending emails.
type SmtpClient struct {
	Email  string
	Dialer Dialer
}

func NewSmtpClient(dialer Dialer) *SmtpClient {
	return &SmtpClient{
		os.Getenv("SMTP_EMAIL"),
		dialer,
	}
}

// BuildEmail builds an email from a template and a context.
func (c SmtpClient) BuildEmail(tmpl string, context map[string]any) ([]byte, error) {
	email, err := template.New("email").Parse(tmpl)
	if err != nil {
		return []byte{}, err
	}
	var data bytes.Buffer
	err = email.Execute(&data, context)
	return data.Bytes(), err
}

// SendEmail sends an email to the given receivers.
func (c SmtpClient) SendEmail(receivers []string, subject string, email []byte) error {
	m := mail.NewMessage()

	m.SetHeader("From", c.Email)
	m.SetHeader("To", receivers...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", string(email))

	return c.Dialer.DialAndSend(m)
}
