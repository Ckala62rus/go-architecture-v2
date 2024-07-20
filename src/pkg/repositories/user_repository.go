package repositories

import "practice/domains"

type UserDB struct {
}

func NewUserDB() *UserDB {
	return &UserDB{}
}

func (u *UserDB) GetUser() domains.User {
	user := domains.User{Name: "Admin", Age: 30}
	return user
}

func (u *UserDB) CreateUser(name string, age int) domains.User {
	user := domains.User{Name: name, Age: age}
	return user
}
