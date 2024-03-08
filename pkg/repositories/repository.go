package repositories

import "practice/domains"

type User interface {
	GetUser() domains.User
	CreateUser(name string, age int) domains.User
}

func NewRepository() *Repository {
	return &Repository{
		User: NewUserDB(),
	}
}

type Repository struct {
	User
}
