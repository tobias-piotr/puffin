package emails

import "log/slog"

type EmailService struct {
	emailRepository EmailRepository
	emailClient     EmailClient
}

func NewEmailService(emailRepository EmailRepository, emailClient EmailClient) EmailService {
	return EmailService{emailRepository: emailRepository, emailClient: emailClient}
}

func (s EmailService) CreateNewTemplate(data *TemplateData) (*Template, error) {
	slog.Info("Creating new template", "name", data.Name)
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
