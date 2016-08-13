package email

import (
	"bytes"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

const (
	tmplInvite = `
You have been invited to join awesome product. {{.InviteURL}}
`
)

var (
	InviteSubject = "You are being invited to awesome product"
	fromEmail     = "aircto@launchyard.com"
)

// START OMIT
type EmailSender interface { // HL
	SendEmail(to []string, from, subject string, body []byte) error
}

var Sender EmailSender // HL

// Amazon SES email sender
type SESEmailSender struct {
	smtpUsername string
	smtpPassword string
	host         string
	port         int
}

// END OMIT

func init() {
	Sender = &SESEmailSender{
		smtpUsername: os.Getenv("SES_SMTP_USERNAME"),
		smtpPassword: os.Getenv("SES_SMTP_PASSWORD"),
		host:         os.Getenv("SES_HOST"),
		port:         587,
	}
}

// START2 OMIT
func (s *SESEmailSender) SendEmail(to []string, from, subject string, body []byte) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", string(body))
	d := gomail.NewPlainDialer(s.host, s.port, s.smtpUsername, s.smtpPassword)
	return d.DialAndSend(m)
}

func SendInvitationEmail(to []string, ctx map[string]string) error {
	body, err := templToBytes("invite", tmplInvite, ctx)
	if err != nil {
		return err
	}
	if err := Sender.SendEmail(to, fromEmail, InviteSubject, body); err != nil { // HL
		return err
	}
	return nil
}

// END2 OMIT
func templToBytes(name, tmpl string, ctx map[string]string) ([]byte, error) {
	var buf bytes.Buffer
	t := template.Must(template.New(name).Parse(tmpl))
	if err := t.Execute(&buf, ctx); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
