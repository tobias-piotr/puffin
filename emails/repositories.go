package emails

type EmailRepository interface {
	CreateNewTemplate(data *TemplateData) (Template, error)
	GetTemplates() ([]Template, error)
}
