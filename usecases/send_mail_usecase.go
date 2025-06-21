package usecases

import (
	"fmt"
	"mail_cobra/domains/user"

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
	return nil
}
