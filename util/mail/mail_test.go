package mail_test

import (
	"known-anchors/util/mail"
	"testing"
)

func TestSendMailCode(t *testing.T) {
	m, err := mail.MailCode("test", "123456", "wizo.o@outlook.com", 5)
	if err != nil {
		t.Error(err)
	}
	if err := mail.SendMail(m); err != nil {
		t.Error(err)
	}
}
