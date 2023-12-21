package mail_test

import (
	"known-anchors/util/mail"
	"testing"
)

func TestSendMailCode(t *testing.T) {
	m, err := mail.MailCode(&mail.MailCodeData{
		Username: "kkk",
		Code:     "123",
		To:       "tto",
		TTL:      1,
	})
	if err != nil {
		t.Error(err)
	}
	if err := mail.SendMail(m); err != nil {
		t.Error(err)
	}
}
