package database

import (
	"puffin/emails"

	"github.com/google/uuid"
)

type EmailRepository struct{}

func (r EmailRepository) CreateNewTemplate(data *emails.TemplateData) (emails.Template, error) {
	return emails.Template{ID: uuid.New(), Name: "test_template", Content: "test content"}, nil
}

func (r EmailRepository) GetTemplates() ([]emails.Template, error) {
	return []emails.Template{
		{
			ID: uuid.New(), Name: "test_template",
			Content: "test content",
		},
		{
			ID: uuid.New(), Name: "test_template",
			Content: "test content",
		},
	}, nil
}
