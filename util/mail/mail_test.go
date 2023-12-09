package mail_test

import (
	"known-anchors/util/mail"
	"testing"
)

func TestSendMailCode(t *testing.T) {
	err := mail.SendMailCode("test", "123456", "wizo.o@outlook.com", 5)
	if err != nil {
		t.Error(err)
	}
}
