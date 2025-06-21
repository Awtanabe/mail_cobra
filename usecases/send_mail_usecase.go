package usecases

import (
	"fmt"
	"log"
	"mail_cobra/domains/user"
	mailer "mail_cobra/sdk/mail"

	"gorm.io/gorm"
)

type SendMailUsecase struct {
	db *gorm.DB
}

func NewSendMailUsecase(db *gorm.DB) *SendMailUsecase {
	return &SendMailUsecase{db: db}
}

func (u *SendMailUsecase) Execute() error {
	r := user.NewUserRepo(u.db)
	users, err := r.ListUser()

	if err != nil {
		return err
	}

	fmt.Print("usecase 実行", users)
	m := mailer.NewMailer(
		"mailhog",
		1025,
		"user@example.com",
		"password",
		"from@example.net",
	)

	for _, u := range *users {

		message := mailer.Message{
			To:      []string{u.Email},
			Subject: "yahoo",
			Body:    "Hello World!",
		}

		if err := m.Send(message); err != nil {
			log.Fatalf("failed to send email: %v", err)
		}
	}

	return nil
}
