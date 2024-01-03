package emails

import (
	"github.com/google/uuid"
)

// TemplateData represents data used for creating new email templates.
type TemplateData struct {
	Name    string `validate:"required,min=3,max=255"`
	Content string `validate:"required"`
}

// Template represents an email template.
type Template struct {
	Id      uuid.UUID
	Name    string
	Content string
}

// EmailData represents data used for sending emails.
type EmailData struct {
	TemplateName string
	To           string
	Subject      string
	Context      map[string]interface{}
}

// Email represents an email with all its data.
type Email struct {
	Id           uuid.UUID
	TemplateName string
	To           string
	Subject      string
	Context      map[string]interface{}
}
