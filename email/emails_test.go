package email

import (
	"bytes"
	"html/template"
	"reflect"
	"testing"
)

// START1 OMIT
type FakeEmailSender struct {
	to   []string
	from string
	body []byte
}

func (f *FakeEmailSender) SendEmail(to []string, from, subject string, body []byte) error {
	f.to, f.from, f.body = to, from, body
	return nil
}

var fakeSender = &FakeEmailSender{} // HL

func init() {
	Sender = fakeSender // HL
}

// END1 OMIT

func check(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func cleanup() {
	fakeSender.to = []string{}
	fakeSender.from = ""
	fakeSender.body = []byte{}
}

// START2 OMIT
func TestSendInviteEmail(t *testing.T) {
	defer cleanup()
	to := []string{"kaviraj@launchyard.com"}
	ctx := map[string]string{"Email": "new-guy@gmail.com"}
	err := SendInvitationEmail(to, ctx)
	check(t, err)
	if !reflect.DeepEqual(fakeSender.to, to) { // HL
		t.Errorf("wanted %q, got %q", to, fakeSender.to)
	}
	if !reflect.DeepEqual(fakeSender.from, fromEmail) { // HL
		t.Errorf("wanted %q, got %q", fromEmail, fakeSender.from)
	}
	tm := template.Must(template.New("test").Parse(tmplInvite))
	var buf bytes.Buffer
	tm.Execute(&buf, ctx)

	if !reflect.DeepEqual(fakeSender.body, buf.Bytes()) { // HL
		t.Errorf("wanted %s, got %s", buf.Bytes(), fakeSender.body)
	}
}

// END2 OMIT
