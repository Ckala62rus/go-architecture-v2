package repositories

import (
	"gorm.io/gorm"
	"practice/domains"
)

type Users interface {
	//CreateUser(user domains.User) (domains.User, error)
	GetUserByName(name string) (domains.User, error)
	GetUserByEmail(email string) (domains.User, error)
	GetById(id int) (domains.User, error)
	GetAllUsers() []domains.User
	DeleteUserById(id int) (bool, error)
	UpdateUser(user domains.User) (domains.User, error)
}

type Authorization interface {
	CreateUser(user domains.User) (domains.User, error)
	//GetUser(email, password string) (domains.User, error)
}

type Repository struct {
	Users
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Users:         NewUserDB(db),
		Authorization: NewAuthGorm(db),
	}
}
