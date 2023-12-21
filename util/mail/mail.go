package mail

import (
	"bytes"
	"encoding/json"
	"known-anchors/config"
	"log"
	"path"
	"text/template"

	"gopkg.in/gomail.v2"
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

func MailToJson(mail *Mail) (string, error) {
	m, err := json.Marshal(mail)
	if err != nil {
		return "", err
	}
	return string(m), nil
}

func JsonToMail(jsonStr string) (*Mail, error) {
	var mail Mail
	err := json.Unmarshal([]byte(jsonStr), &mail)
	if err != nil {
		return nil, err
	}
	return &mail, nil
}

type MailCodeData struct {
	Username string
	Code     string
	To       string
	TTL      int
}

func MailCodeToJson(username, code, to string, ttl int) (string, error) {
	m := MailCodeData{
		Username: username,
		Code:     code,
		To:       to,
		TTL:      ttl,
	}
	mail, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(mail), nil
}

func JsonToMailCode(jsonStr string) (*MailCodeData, error) {
	var mail MailCodeData
	err := json.Unmarshal([]byte(jsonStr), &mail)
	if err != nil {
		return nil, err
	}
	return &mail, nil
}

func MailCode(m *MailCodeData) (*Mail, error) {
	templatePath := path.Join("template", "views", "mailcode.tmpl")
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, map[string]interface{}{
		"gencode":  m.Code,
		"username": m.Username,
		"ttl":      m.TTL,
	}); err != nil {
		return nil, err
	}
	mail := &Mail{
		From:    config.Conf.Mail.UserName,
		To:      m.To,
		Subject: "Known-Anchor 验证码",
		Body:    tpl.String(),
	}
	return mail, nil
}

func SendMail(mail *Mail) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", mail.From)
	m.SetHeader("To", mail.To)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.Body)
	if err = mailer.DialAndSend(m); err != nil {
		log.Println(err)
	}
	return err
}
