package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"practice/pkg/dto"
	"strconv"
)

type GetUser struct {
	Name string `uri:"name" binding:"required"`
}

// Hello
// @Summary      Testing route
// @Description  get test json response
// @Param name   path string true "username"
// @Tags         Testing
// @Accept       json
// @Produce      json
// @Router       /hello/{name} [get]
func (h *Handler) Hello(c *gin.Context) {
	var getUser GetUser
	err := c.ShouldBindUri(&getUser)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSONP(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "Hello handler",
		Data:    getUser.Name,
	})
}

// GetAllUsers
// @Summary      Get all users
// @Description  return all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  StatusResponse
// @Router       /users/ [get]
// @Security Authorization
func (h *Handler) GetAllUsers(c *gin.Context) {
	users := h.services.Users.GetAllUsers()
	//var usersDTO []dto.AllUsersOutDTO

	//for _, user := range users {
	//	dtoUserMap := dto.AllUsersOutDTO{
	//		Id:        user.Id,
	//		Name:      user.Name,
	//		Email:     user.Email,
	//		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	//		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	//	}
	//	usersDTO = append(usersDTO, dtoUserMap)
	//}

	usersDTO := dto.MapAllUser(users)

	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "all users",
		Data:    usersDTO,
	})
}

// GetUserByName
// @Summary      Get user by Name
// @Description  get user by Name
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        name path string  true "User name"
// @Success      200  {object}  StatusResponse
// @Router       /users/user/{name} [get]
// @Security Authorization
func (h *Handler) GetUserByName(c *gin.Context) {
	name := c.Param("name")
	user, err := h.services.Users.GetUserByName(name)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	userMap := dto.MapSingleUser(user)
	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "one user",
		Data:    userMap,
	})
}

// GetById
// @Summary      Get user by ID
// @Description  get user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  StatusResponse
// @Router       /users/{id} [get]
// @Security Authorization
func (h *Handler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	user, err := h.services.Users.GetById(id)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	userMap := dto.MapSingleUser(user)

	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "one user",
		Data:    userMap,
	})
}
