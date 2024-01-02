package emails

type EmailClient interface {
	SendEmail() error
}
