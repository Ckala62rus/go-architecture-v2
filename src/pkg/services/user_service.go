package services

import (
	"practice/domains"
	"practice/pkg/repositories"
)

type UserService struct {
	repo repositories.Users
}

func NewUserService(rep repositories.Users) *UserService {
	return &UserService{repo: rep}
}

func (u *UserService) GetUserByName(name string) (domains.User, error) {
	return u.repo.GetUserByName(name)
}

func (u *UserService) GetById(id int) (domains.User, error) {
	return u.repo.GetById(id)
}

func (u *UserService) GetAllUsers() []domains.User {
	return u.repo.GetAllUsers()
}

func (u *UserService) CreateUser(user domains.User) (domains.User, error) {
	return u.repo.CreateUser(user)
}

func (u *UserService) DeleteUserById(id int) (bool, error) {
	return u.repo.DeleteUserById(id)
}

func (u *UserService) UpdateUser(userRequest domains.User) (domains.User, error) {
	return u.repo.UpdateUser(userRequest)
}
