package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"practice/pkg/services"
)

type Handler struct {
	services *services.Service
	log      *slog.Logger
}

func NewHandler(services *services.Service, log *slog.Logger) *Handler {
	return &Handler{
		services: services,
		log:      log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/hello/:name", h.Hello)
	}

	return router
}
