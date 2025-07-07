package handlers

import "github.com/gin-gonic/gin"

// ErrorResponse представляет стандартный ответ с ошибкой
type ErrorResponse struct {
	Message string  `json:"message" example:"Пользователь не найден" description:"Описание ошибки"`
	Status  bool    `json:"status" example:"false" description:"Статус запроса (всегда false для ошибок)"`
	Data    *string `json:"data" example:"null" description:"Дополнительные данные об ошибке (обычно null)"`
}

// StatusResponse представляет стандартный успешный ответ
type StatusResponse struct {
	Status  bool        `json:"status" example:"true" description:"Статус запроса (true для успешных операций)"`
	Message string      `json:"message" example:"Операция выполнена успешно" description:"Сообщение о результате операции"`
	Data    interface{} `json:"data" swaggerignore:"true" description:"Данные ответа (пользователь, список пользователей, токен и т.д.)"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(
		statusCode,
		ErrorResponse{
			Message: message,
			Status:  false,
			Data:    nil,
		},
	)
}
