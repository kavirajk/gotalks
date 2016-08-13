package email

import (
	"bytes"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

// START1 OMIT
const (
	tmplInvite = `
You have been invited to join awesome product. {{.InviteURL}}
`
)

// END1 OMIT

// START2 OMIT
var (
	InviteSubject = "You are being invited to awesome product"
	fromEmail     = "aircto@launchyard.com"
	config        *SESEmailConfig
)

// Amazon SES email config
type SESEmailConfig struct {
	smtpUsername string
	smtpPassword string
	host         string
	port         int
}

func sendEmail(to []string, from, subject string, body []byte) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", string(body))
	d := gomail.NewPlainDialer(config.host, config.port, config.smtpUsername, config.smtpPassword)
	return d.DialAndSend(m)
}

// END2 OMIT
func init() {
	config = &SESEmailConfig{
		smtpUsername: os.Getenv("SES_SMTP_USERNAME"),
		smtpPassword: os.Getenv("SES_SMTP_PASSWORD"),
		host:         os.Getenv("SES_HOST"),
		port:         587,
	}
}

func SendInvitationEmail(to []string, ctx map[string]string) error {
	body, err := templToBytes("invite", tmplInvite, ctx)
	if err != nil {
		return err
	}
	if err := sendEmail(to, fromEmail, InviteSubject, body); err != nil { // HL
		return err
	}
	return nil
}

func templToBytes(name, tmpl string, ctx map[string]string) ([]byte, error) {
	var buf bytes.Buffer
	t := template.Must(template.New(name).Parse(tmpl))
	if err := t.Execute(&buf, ctx); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
