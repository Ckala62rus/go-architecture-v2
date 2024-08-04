package repositories

import (
	"gorm.io/gorm"
	"practice/domains"
)

type Users interface {
	CreateUser(user domains.User) (domains.User, error)
	GetUserByName(name string) (domains.User, error)
	GetById(id int) (domains.User, error)
	GetAllUsers() []domains.User
	DeleteUserById(id int) (bool, error)
	UpdateUser(user domains.User) (domains.User, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Users: NewUserDB(db),
	}
}

type Repository struct {
	Users
}
