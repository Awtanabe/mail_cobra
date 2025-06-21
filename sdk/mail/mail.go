package mailer

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Mailer struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	Auth     smtp.Auth
	SmtpAddr string
}

type Message struct {
	To      []string
	Subject string
	Body    string
}

func NewMailer(host string, port int, username, password, from string) *Mailer {
	auth := smtp.CRAMMD5Auth(username, password)
	return &Mailer{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
		Auth:     auth,
		SmtpAddr: fmt.Sprintf("%s:%d", host, port),
	}
}

func (m *Mailer) Send(msg Message) error {
	message := m.BuildMessage(msg)
	return smtp.SendMail(m.SmtpAddr, m.Auth, m.From, msg.To, message)
}

func (m *Mailer) BuildMessage(msg Message) []byte {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n"

	return []byte(fmt.Sprintf("To: %s\nSubject: %s\n%s\n\n%s",
		strings.Join(msg.To, ","),
		msg.Subject,
		mime,
		msg.Body,
	))
}
