package emails

type EmailRepository interface {
	CreateNewTemplate() (Email, error)
	GetTemplates() ([]Template, error)
}
