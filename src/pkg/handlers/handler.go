package handlers

import (
	"log/slog"
	"net/http"
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

	// redirect on swagger ui dashboard
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	api := router.Group("/api")
	{
		api.GET("/hello/:name", gin.BasicAuth(gin.Accounts{
			"owner": "123123",
		}), h.Hello)

		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.SignUp)
			auth.POST("/sign-in", h.SignIn)
			auth.GET("/me", h.userIdentity, h.Me)
			auth.POST("/logout", h.userIdentity, h.Logout)
		}

		users := api.Group("/users", h.userIdentity)
		{
			users.GET("/", h.GetAllUsers)
			users.GET("/user/:name", h.GetUserByName)
			users.GET("/:id", h.GetById)
			//users.POST("/", h.CreateUser)
			users.DELETE(":id", h.DeleteUserById)
			users.PUT(":id", h.UpdateUser)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
