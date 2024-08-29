package services

import (
	"practice/domains"
	"practice/pkg/repositories"
)

type Users interface {
	//CreateUser(user domains.User) (domains.User, error)
	GetUserByName(name string) (domains.User, error)
	GetById(id int) (domains.User, error)
	GetAllUsers() []domains.User
	DeleteUserById(id int) (bool, error)
	UpdateUser(userRequest domains.User) (domains.User, error)
}

type Authorization interface {
	CreateUser(user domains.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Users
	Authorization
}

func NewService(rep *repositories.Repository) *Service {
	return &Service{
		Users:         NewUserService(rep.Users),
		Authorization: NewAuthService(rep.Authorization, rep.Users),
	}
}
