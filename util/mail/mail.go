package mail

import (
	"bytes"
	"context"
	"known-anchors/config"
	"path"
	"text/template"

	"gopkg.in/gomail.v2"
	"gorm.io/gorm/logger"
)

var (
	mailer = gomail.NewDialer(config.Conf.Mail.Host, config.Conf.Mail.Port, config.Conf.Mail.UserName, config.Conf.Mail.Password)
)

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
}

func SendMailCode(username, code, to string, ttl int) error {
	templatePath := path.Join("template", "views", "mailcode.tmpl")
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, map[string]interface{}{
		"gencode":  code,
		"username": username,
		"ttl":      ttl,
	}); err != nil {
		return err
	}
	mail := &Mail{
		From:    config.Conf.Mail.UserName,
		To:      to,
		Subject: "Known-Anchor 验证码",
		Body:    tpl.String(),
	}
	return SendMail(mail)
}

func SendMail(mail *Mail) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", mail.From)
	m.SetHeader("To", mail.To)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.Body)
	if err = mailer.DialAndSend(m); err != nil {
		logger.Default.Error(context.Background(), err.Error())
	}
	return err
}
