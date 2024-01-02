package emails

type EmailService struct {
	emailRepository EmailRepository
	emailClient     EmailClient
}

func (s EmailService) CreateNewTemplate() (Template, error) {
	return Template{}, nil // TODO
}

func (s EmailService) GetTemplates() ([]Template, error) {
	return []Template{}, nil // TODO
}

func (s EmailService) SendEmail() error {
	return s.emailClient.SendEmail()
}
