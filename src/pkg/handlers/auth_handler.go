package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"practice/domains"
	"practice/pkg/dto"
)

// SignUp
// @Summary      Authentication in system
// @Description  return id created user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param input body dto.CreateAuthUser true "credentials"
// @Success      200  {object}  StatusResponse
// @Router       /auth/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
	var input dto.CreateAuthUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(domains.User{
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "User success created",
		Data:    id,
	})
}

// SignIn
// @Summary Login
// @Description login and return authorization bearer token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body dto.SignInInput true "credentials"
// @Success	200  {object}  StatusResponse
// @Router /auth/sign-in [post]
func (h *Handler) SignIn(c *gin.Context) {
	var input dto.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusOK, StatusResponse{
			Status:  false,
			Message: "authentication failed",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	ctx := context.Background()
	RedisCache.SetToken(ctx, token)

	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "authentication success!",
		Data: map[string]interface{}{
			"token": "Bearer " + token,
		},
	})
}

// Me
// @Summary 	 User information
// @Tags         auth
// @Description  get authorization user information by id
// @Accept       json
// @Produce      json
// @Success      200  {object}  StatusResponse
// @Failure      401 {object} ErrorResponse
// @Router       /auth/me [get]
// @Security Authorization
func (h *Handler) Me(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	user, err := h.services.Users.GetById(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userDTO := dto.MapSingleUser(user)

	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "images was updated",
		Data:    userDTO,
	})
}
