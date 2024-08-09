package repositories

import (
	"gorm.io/gorm"
	"practice/domains"
)

type AuthGorm struct {
	db *gorm.DB
}

func NewAuthGorm(db *gorm.DB) *AuthGorm {
	return &AuthGorm{db: db}
}

func (u *AuthGorm) CreateUser(user domains.User) (domains.User, error) {
	result := u.db.Create(&user)
	return user, result.Error
}
