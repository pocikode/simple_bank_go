package mail

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type EmailSender interface {
	Send(
		subject,
		content string,
		to,
		cc,
		bcc,
		attachFiles []string,
	) error
}

type Email struct {
	name             string
	fromEmailAddress string
	host             string
	port             int
	username         string
	password         string
}

func NewEmailSender(name, fromEmailAddress, host string, port int, username, password string) EmailSender {
	return &Email{
		name:             name,
		fromEmailAddress: fromEmailAddress,
		host:             host,
		port:             port,
		username:         username,
		password:         password,
	}
}

func (sender *Email) Send(subject, content string, to, cc, bcc, attachFiles []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		if _, err := e.AttachFile(f); err != nil {
			return fmt.Errorf("failed to attach file %s: %w", f, err)
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.username, sender.password, sender.host)
	return e.Send(fmt.Sprintf("%s:%d", sender.host, sender.port), smtpAuth)
}
