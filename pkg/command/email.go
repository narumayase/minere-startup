package command

import (
	"crypto/tls"
	gomail "gopkg.in/mail.v2"
	"log"
	"minere-startup/pkg/config"
)

type Email struct {
	Config *config.Configuration
}

func NewEmailCommand(c *config.Configuration) *Email {
	return &Email{Config: c}
}

func (e *Email) Send(subject, body string) {
	m := gomail.NewMessage()

	m.SetHeader("From", e.Config.Email.From)
	m.SetHeader("To", e.Config.Email.To)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(e.Config.Email.Smtp, e.Config.Email.SmtpPort, e.Config.Email.From, e.Config.Email.PasswordFrom)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("email successfully send!!")
}
