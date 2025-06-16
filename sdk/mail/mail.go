package mailer

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Mailer struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type Message struct {
	To      []string
	Subject string
	Body    string
}

func NewMailer(host string, port int, username, password, from string) *Mailer {
	return &Mailer{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

func (m *Mailer) Send(msg Message) error {
	smtpAddr := fmt.Sprintf("%s:%d", m.Host, m.Port)

	log.Print("smtpAddr", smtpAddr)

	auth := smtp.CRAMMD5Auth(m.Username, m.Password)

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n"
	message := []byte(fmt.Sprintf("To: %s\nSubject: %s\n%s\n\n%s",
		strings.Join(msg.To, ","),
		msg.Subject,
		mime,
		msg.Body,
	))

	return smtp.SendMail(smtpAddr, auth, m.From, msg.To, message)
}
