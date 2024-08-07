package handlers

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Message string      `json:"message" default:"Explaining of error description"`
	Status  bool        `json:"status" default:"False"`
	Data    interface{} `json:"data"`
}

type StatusResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(
		statusCode,
		ErrorResponse{
			message,
			false,
			nil,
		},
	)
}
