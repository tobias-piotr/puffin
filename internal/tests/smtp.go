package tests

import (
	"github.com/go-mail/mail"
)

type Dialer interface {
	DialAndSend(m ...*mail.Message) error
}
type DummyDialer struct {
	Emails []*mail.Message
}

func (d *DummyDialer) DialAndSend(m ...*mail.Message) error {
	d.Emails = append(d.Emails, m...)
	return nil
}
