package example

type OldEmailSender interface {
	Send(from, to, subject, body string) error
}

type OldEmailSvc struct{}

func (s OldEmailSvc) Send(from, to, subject, body string) error {
	return nil
}

type NewEmailSender interface {
	Send() error
}

type NewEmailService struct {
	From, To, Subject, Body string
}

func (s NewEmailService) Send() error {
	return nil
}

type OldEmailServiceAdapter struct {
	From, To, Subject, Body string
	OldEmailService         OldEmailSender
}

func (a OldEmailServiceAdapter) Send() error {
	return a.OldEmailService.Send(a.From, a.To, a.Subject, a.Body)
}

type EmailClient struct {
	Mail NewEmailSender
}

func (e EmailClient) SendEmail(From, To, Subject, Body string) error {
	return e.Mail.Send()
}