package emails

import (
	"fmt"
	"net/smtp"
	"os"
)

type SmtpClient struct {
	Email    string
	Password string
	Host     string
	Port     int
}

func NewSmtpClient() *SmtpClient {
	return &SmtpClient{
		os.Getenv("SMTP_EMAIL"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("SMTP_HOST"),
		587,
	}
}

func (c SmtpClient) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c SmtpClient) SendEmail(receivers []string, message []byte) error {
	auth := smtp.PlainAuth("", c.Email, c.Password, c.Host)
	return smtp.SendMail(c.GetAddr(), auth, c.Email, receivers, message)
}
