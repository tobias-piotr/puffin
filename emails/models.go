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

// EmailData represents data used for sending emails.
type EmailData struct {
	TemplateName string         `json:"template_name" validate:"required" db:"template_name"`
	Recipients   []string       `json:"recipients" validate:"required"`
	Subject      string         `json:"subject" validate:"required"`
	Context      map[string]any `json:"context"`
	// TODO: Validate email addresses
}

// Email represents an email with all its data.
type Email struct {
	ID           uuid.UUID      `json:"id"`
	CreatedAt    time.Time      `json:"created_at" db:"created_at"`
	TemplateName string         `json:"template_name" db:"template_name"`
	Recipients   []string       `json:"recipients"`
	Subject      string         `json:"subject"`
	Context      map[string]any `json:"context"`
}
