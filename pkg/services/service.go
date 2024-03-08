package services

import (
	"practice/domains"
	"practice/pkg/repositories"
)

type Users interface {
	GetUser() domains.User
	CreateUser(name string, age int) domains.User
}

type Service struct {
	Users
}

func NewService(rep *repositories.Repository) *Service {
	return &Service{
		Users: NewUserService(rep.User),
	}
}
