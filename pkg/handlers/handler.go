package handlers

import (
	"github.com/gin-gonic/gin"
	"practice/pkg/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/hello/:name", h.Hello)
	}

	return router
}
