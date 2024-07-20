package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUser struct {
	Name string `uri:"name" binding:"required"`
}

func (h *Handler) Hello(c *gin.Context) {
	user := h.services.GetUser()

	h.log.Debug("handler hello")

	var getUser GetUser
	err := c.ShouldBindUri(&getUser)
	if err != nil {
		fmt.Println(err.Error())
	}

	h.log.Debug(getUser.Name)

	fmt.Println(getUser.Name)

	c.JSONP(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "Hello handler",
		Data:    user,
	})
}
