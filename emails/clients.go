package emails

type EmailClient interface {
	BuildEmail(tmpl string, context map[string]any) ([]byte, error)
	SendEmail(receivers []string, subject string, email []byte) error
}
