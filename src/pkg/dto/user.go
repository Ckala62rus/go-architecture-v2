package dto

import (
	"practice/domains"
	"time"
)

// CreateUserInDTO represents the input data for creating a new user
type CreateUserInDTO struct {
	Name     string `json:"name" binding:"required,min=3,max=50" example:"admin" format:"string" description:"Имя пользователя (от 3 до 50 символов)"`
	Email    string `json:"email" binding:"required,email" example:"admin@mail.ru" format:"email" description:"Email адрес пользователя"`
	Password string `json:"password" binding:"required,min=6" example:"123123" format:"password" description:"Пароль пользователя (минимум 6 символов)"`
}

// UpdateUserInDTO represents the input data for updating a user
type UpdateUserInDTO struct {
	Name string `json:"name" binding:"required,min=3,max=50" example:"new_admin" format:"string" description:"Новое имя пользователя (от 3 до 50 символов)"`
}

// UserOutDTO represents the output data for a user
type UserOutDTO struct {
	Id        int    `json:"id" example:"1" description:"Уникальный идентификатор пользователя"`
	Name      string `json:"name" example:"admin" description:"Имя пользователя"`
	Email     string `json:"email" example:"admin@mail.ru" description:"Email адрес пользователя"`
	CreatedAt string `json:"created_at" example:"2024-01-01 12:00:00" description:"Дата и время создания пользователя"`
	UpdatedAt string `json:"updated_at" example:"2024-01-01 12:00:00" description:"Дата и время последнего обновления пользователя"`
}

func MapAllUser(users []domains.User) []UserOutDTO {
	var usersDTO []UserOutDTO

	for _, user := range users {
		dtoUserMap := UserOutDTO{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		usersDTO = append(usersDTO, dtoUserMap)
	}

	return usersDTO
}

func MapSingleUser(user domains.User) UserOutDTO {

	usersDTO := UserOutDTO{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.DateTime), // "2006-01-02 15:04:05"
		UpdatedAt: user.UpdatedAt.Format(time.DateTime),
	}

	return usersDTO
}

// CreateAuthUser represents the input data for user registration
type CreateAuthUser struct {
	Email    string `json:"email" binding:"required,email" example:"agr.akyla@mail.ru" format:"email" description:"Email адрес для регистрации"`
	Password string `json:"password" binding:"required,min=6" example:"123123" format:"password" description:"Пароль для регистрации (минимум 6 символов)"`
}

// SignInInput represents the input data for user authentication
type SignInInput struct {
	Email    string `json:"email" binding:"required,email" example:"agr.akyla@mail.ru" format:"email" description:"Email адрес для входа"`
	Password string `json:"password" binding:"required" example:"123123" format:"password" description:"Пароль для входа"`
}
