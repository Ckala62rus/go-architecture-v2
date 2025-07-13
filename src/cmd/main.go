package main

import (
	"fmt"
	"practice/domains"
	"practice/internal/logger"
	handlers "practice/pkg/handlers"
	"practice/pkg/repositories"
	"practice/pkg/services"
	"practice/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title           Go Template API
// @version         1.0.0
// @description     RESTful API сервер на Go с использованием Gin, PostgreSQL, Redis, JWT аутентификацией и чистой архитектурой
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://github.com/your-username/practice
// @contact.email  support@yourcompany.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:5000
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT токен с префиксом "Bearer " (например: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...")

// @externalDocs.description  GitHub Repository
// @externalDocs.url          https://github.com/your-username/practice
func main() {
	cfg := utils.MainConfig
	//fmt.Printf("%+v", cfg) // todo need delete!

	// todo init logger: slog
	log := logger.MainLogger
	//log.Info("Start application!")
	log.Info("***********************************************************")

	// todo init storage: gorm
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.User,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.Db,
		cfg.DatabaseConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if cfg.Env == "local" || cfg.Env == "dev" {
		db.Debug()
	}
	if err != nil {
		panic(err)
	}

	// Run migration initialize
	utils.AutoMigrateInitialize(db, log, cfg.Env)

	// todo init router: gin-gonic router
	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	rep := repositories.NewRepository(db)
	service := services.NewService(rep)
	handlerCollection := handlers.NewHandler(service, log)

	// todo server: standard golang server
	out := fmt.Sprintf("server started in localhost:%s", cfg.HttpServer.Port)
	log.Info(out)

	srv := new(domains.Server)
	if err := srv.Run(cfg.HttpServer.Port, handlerCollection.InitRoutes()); err != nil {
		log.Debug("error occurred while running http server", "error", err.Error())
	}
}
