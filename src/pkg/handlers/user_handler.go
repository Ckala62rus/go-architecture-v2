package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
