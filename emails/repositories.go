package emails

type (
	EmailFilters    map[string]any
	EmailRepository interface {
		CreateNewTemplate(data *TemplateData) (*Template, error)
		GetTemplates() ([]Template, error)
		FilterTemplates(filters EmailFilters) ([]Template, error)
		SaveEmail(data *EmailData) (*Email, error)
		GetEmails() ([]Email, error)
	}
)
