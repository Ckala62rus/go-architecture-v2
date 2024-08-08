package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"practice/domains"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) GetUserByName(name string) (domains.User, error) {
	user := domains.User{}
	u.db.Debug().Where("name = ?", name).First(&user)
	if user.Id == 0 {
		mistake := fmt.Sprintf("user with name:%s not found", name)
		return user, errors.New(mistake)
	}
	return user, nil
}

func (u *UserDB) GetUserByEmail(email string) (domains.User, error) {
	user := domains.User{}
	u.db.Where(map[string]interface{}{"email": email}).Find(&user)
	if user.Id == 0 {
		msg := fmt.Sprintf("User with email %s not found", email)
		return user, errors.New(msg)
	}
	return user, nil
}

func (u *UserDB) GetById(id int) (domains.User, error) {
	var user domains.User
	u.db.First(&user, id)
	if user.Id == 0 {
		mistake := fmt.Sprintf("user with id:%d not found", id)
		return user, errors.New(mistake)
	}
	return user, nil
}

func (u *UserDB) GetAllUsers() []domains.User {
	var users []domains.User
	u.db.Debug().Order("id desc").Find(&users)
	return users
}

//func (u *UserDB) CreateUser(user domains.User) (domains.User, error) {
//	result := u.db.Create(&user)
//	return user, result.Error
//}

func (u *UserDB) DeleteUserById(id int) (bool, error) {
	res := u.db.Delete(&domains.User{}, id)
	intDelete := res.RowsAffected
	err := res.Error

	if err != nil || intDelete == 0 {
		mistake := fmt.Sprintf("can't delete user with id:%d", id)
		return false, errors.New(mistake)
	}

	return true, nil
}

func (u *UserDB) UpdateUser(userRequest domains.User) (domains.User, error) {
	var user domains.User
	u.db.Debug().First(&user, userRequest.Id)
	if user.Id == 0 {
		mistake := fmt.Sprintf("user not found with id:%d", userRequest.Id)
		return user, errors.New(mistake)
	}

	if user.Name != userRequest.Name && len(userRequest.Name) > 0 {
		user.Name = userRequest.Name
	}

	u.db.Save(user)
	return user, nil
}
