package user

import (
	"gorm.io/gorm"
)

type UserRepo interface {
	ListUser() (*[]User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) ListUser() (*[]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}
