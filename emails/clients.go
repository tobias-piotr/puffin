package emails

type EmailClient interface {
	SendEmail(data *EmailData) error
}
