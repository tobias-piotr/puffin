package emails

import (
	"fmt"
	"log/slog"
	"net/http"

	"puffin/libs/api"
)

type EmailService struct {
	emailRepository EmailRepository
	emailClient     EmailClient
}

func NewEmailService(emailRepository EmailRepository, emailClient EmailClient) EmailService {
	return EmailService{emailRepository: emailRepository, emailClient: emailClient}
}

func (s EmailService) CreateNewTemplate(data *TemplateData) (*Template, error) {
	slog.Info("Creating new template", "name", data.Name)
	templates, err := s.emailRepository.FilterTemplates(EmailFilters{"name": data.Name})
	if err != nil {
		return &Template{}, err
	}
	if len(templates) > 0 {
		return &Template{}, api.NewAPIError(
			http.StatusConflict,
			fmt.Sprintf("Template with name %s already exists", data.Name),
			nil,
		)
	}
	return s.emailRepository.CreateNewTemplate(data)
}

func (s EmailService) GetTemplates() ([]Template, error) {
	slog.Info("Getting templates")
	return s.emailRepository.GetTemplates()
}

func (s EmailService) SendEmail(data *EmailData) error {
	slog.Info("Sending email", "template", data.TemplateName)
	return s.emailClient.SendEmail(data)
}
