package emails

import (
	"time"

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
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
}

// EmailContext is a map of variables used in email templates.
type EmailContext map[string]any

// EmailData represents data used for sending emails.
type EmailData struct {
	TemplateName string       `json:"template_name" validate:"required"`
	To           []string     `json:"to" validate:"required"`
	Subject      string       `json:"subject" validate:"required"`
	Context      EmailContext `json:"context"`
	// TODO: Validate email addresses
}

// Email represents an email with all its data.
type Email struct {
	ID           uuid.UUID    `json:"id"`
	TemplateName string       `json:"template_name"`
	To           string       `json:"to"`
	Subject      string       `json:"subject"`
	Context      EmailContext `json:"context"`
}
