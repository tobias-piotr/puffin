package emails

import (
	"github.com/google/uuid"
)

// TemplateData represents data used for creating new email templates.
type TemplateData struct {
	Name    string `validate:"required,min=3,max=255"`
	Content string `validate:"required"`
	// TODO: Add support for attachments
}

// Template represents an email template.
type Template struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Content string    `json:"content"`
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
	ID           uuid.UUID              `json:"id"`
	TemplateName string                 `json:"template_name"`
	To           string                 `json:"to"`
	Subject      string                 `json:"subject"`
	Context      map[string]interface{} `json:"context"`
}
