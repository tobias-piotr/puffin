package database

import (
	"puffin/emails"
	"puffin/libs/database"

	"github.com/lib/pq"
)

// EmailData is the data needed to create an email.
type EmailData struct {
	emails.EmailData
	Recipients pq.StringArray
	Context    database.JSON
}

// Email represents an email in the database.
type Email struct {
	emails.Email
	Recipients pq.StringArray
	Context    database.JSON
}

// ToEmail converts an Email to the domain object.
func (e *Email) ToEmail() *emails.Email {
	e.Email.Recipients = []string(e.Recipients)
	e.Email.Context = map[string]any(e.Context)
	return &e.Email
}
