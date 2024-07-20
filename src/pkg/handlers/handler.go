package handlers

import (
	"log/slog"
	"practice/pkg/services"

	_ "practice/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files

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

// @Summary      Testing route
// @Description  get test json response
// @Param name   path string true "user name"
// @Tags         Testing
// @Accept       json
// @Produce      json
// @Router       /hello/{name} [get]
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/hello/:name", h.Hello)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
