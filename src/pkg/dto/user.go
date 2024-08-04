package dto

import (
	"practice/domains"
)

type CreateUserInDTO struct {
	Name     string `json:"name" example:"admin"`
	Email    string `json:"email" example:"admin@mail.ru"`
	Password string `json:"password" example:"123123"`
}

type UpdateUserInDTO struct {
	Name string `json:"name"`
}

type UserOutDTO struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return usersDTO
}
