package services

import (
	"practice/domains"
	"practice/pkg/repositories"
)

type UserService struct {
	rep repositories.User
}

func NewUserService(rep repositories.User) *UserService {
	return &UserService{rep: rep}
}

func (rep UserService) GetUser() domains.User {
	return rep.rep.GetUser()
}

func (rep UserService) CreateUser(name string, age int) domains.User {
	return rep.rep.CreateUser(name, age)
}
