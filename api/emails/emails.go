package emails

import "net/http"

type EmailSender interface {
	SendEmailConfirmation() error
	SendInvitation() error
	SendPasswordReset() error
}

type OneSignalEmailSender struct {
	baseURL  string
	appToken string
	client   http.Client
}

func (es OneSignalEmailSender) SendEmailConfirmation() error {
	templateID := ""
	return es.doOneSignalRequest(templateID, nil)
}

func (es OneSignalEmailSender) SendInvitation() error {
	templateID := ""
	return es.doOneSignalRequest(templateID, nil)
}

func (es OneSignalEmailSender) SendPasswordReset() error {
	templateID := ""
	return es.doOneSignalRequest(templateID, nil)
}

func (es OneSignalEmailSender) doOneSignalRequest(templateID string, data map[string]any) error {
	return nil
}
