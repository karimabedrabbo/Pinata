package managers

import (
	"github.com/karimabedrabbo/eyo/api/apputils"
	"github.com/karimabedrabbo/eyo/api/constants"
	"github.com/matcornic/hermes/v2"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Mail struct {
	hermesClient *hermes.Hermes
}

var m *Mail

func SetupMail() *Mail {
	return &Mail{
		hermesClient: &hermes.Hermes{
			Product: hermes.Product{
				Name: apputils.GetAppName(),
				Link: apputils.GetAppUrl(),
			},
		},
	}
}

func InitMail() {
	m = SetupMail()
}

func GetMail() *Mail {
	return m
}

func (m *Mail) getBodyFromTemplate(template hermes.Email) (string, error) {
	return m.hermesClient.GenerateHTML(template)
}

func (m *Mail) getFeedbackTemplate(name string, email string, digitCode string) {

}

func (m *Mail) getReportTemplate(reason string, messageHistoryPrettyJson string, noActionResolveUrl string, banOffenderUrl string, banSenderUrl string) hermes.Email {
	return hermes.Email{
		Body:hermes.Body{
			Title: "[" + apputils.GetAppName() + " REPORT] " + reason,
			Intros: []string{
				"Message history:",
				messageHistoryPrettyJson,
			},
			Actions: []hermes.Action{
				{
					Button: hermes.Button{
						Color: "#0276FD",
						Text:  "NO ACTION RESOLVE",
						Link:  noActionResolveUrl,
					},
				},
				{
					Button: hermes.Button{
						Color: "#FF4500",
						Text:  "BAN OFFENDER",
						Link:  banOffenderUrl,
					},
				},
				{
					Button: hermes.Button{
						Color: "#9400D3",
						Text:  "BAN SENDER",
						Link:  banSenderUrl,
					},
				},
			},
		},
	}
}

func (m *Mail) getVerifyEmailTemplate(name string, verifyUrl string) hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name: name,
			Intros: []string{
				"Welcome to " + apputils.GetAppName() + "! Good to have you here.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Verify your account by using this link.",
					Button: hermes.Button{
						Color: "#0276FD",
						Text:  "Verify Account",
						Link:  verifyUrl,
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
}

func (m *Mail) getPasswordResetTemplate(name string, resetUrl string) hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name: name,
			Intros: []string{
				"We heard you need a password reset for " + apputils.GetAppName() + ". We're here to help!",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Use this link to reset your password. If you did not request a password reset please simply ignore this email :).",
					Button: hermes.Button{
						Color: "#0276FD",
						Text:  "Reset Password",
						Link:  resetUrl,
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
}

func (m *Mail) SendReport(reason string, messageHistoryPrettyJson string, noActionResolveUrl string, banOffenderUrl string, banSenderUrl string) error {
	return nil
}

func (m *Mail) SendFeedback(name string, email string) error {
	return nil
}

func (m *Mail) SendVerifyEmail(name string, email string, verifyUrl string) error {
	var err error

	emailTemplate := m.getVerifyEmailTemplate(name, verifyUrl)
	emailBody, err := m.getBodyFromTemplate(emailTemplate)
	if err != nil {
		return err
	}

	from := mail.NewEmail(apputils.GetAppName(), k.AppSupportEmail)
	to := mail.NewEmail(name, email)
	message := mail.NewSingleEmail(from, "Verify Email", to, emailBody, emailBody)
	client := sendgrid.NewSendClient(apputils.GetSendgridApiKey())
	_, err = client.Send(message)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mail) SendPasswordReset(name string, email string, resetUrl string) error {
	var err error

	emailTemplate := m.getPasswordResetTemplate(name, resetUrl)
	emailBody, err := m.getBodyFromTemplate(emailTemplate)
	if err != nil {
		return err
	}
	from := mail.NewEmail(apputils.GetAppName(), k.AppSupportEmail)
	to := mail.NewEmail(name, email)
	message := mail.NewSingleEmail(from, "Reset Password", to, emailBody, emailBody)
	client := sendgrid.NewSendClient(apputils.GetSendgridApiKey())
	_, err = client.Send(message)
	if err != nil {
		return err
	}

	return nil
}
