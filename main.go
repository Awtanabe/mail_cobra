package main

import (
	"log"
	"mail_cobra/cmd"
	mailer "mail_cobra/sdk/mail"
)

func main() {

	m := mailer.NewMailer(
		"mailhog",
		1025,
		"user@example.com",
		"password",
		"from@example.net",
	)
	message := mailer.Message{
		To:      []string{"receiver@example.com"},
		Subject: "hello",
		Body:    "Hello World!",
	}

	if err := m.Send(message); err != nil {
		log.Fatalf("failed to send email: %v", err)
	}

	cmd.Execute()
}
