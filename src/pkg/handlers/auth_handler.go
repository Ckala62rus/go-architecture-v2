package handlers

import (
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
