package main

import (
	"fmt"
	"os"
	"path/filepath"
	"practice/domains"
	config "practice/internal/config"
	"practice/internal/logger"
	handlers "practice/pkg/handlers"
	"practice/pkg/repositories"
	"practice/pkg/services"
	"practice/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5000
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// todo init config: cleanenv
	projectDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cfg := config.MustLoad(filepath.Join(projectDir, "config", "local.yml"))
	//fmt.Printf("%+v", cfg) // todo need delete!

	// todo init logger: slog
	log := logger.SetupNewLogger(cfg.Env, filepath.Join(projectDir, "logs", "logs.log"))
	log.Info("Start application!")

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
	utils.AutoMigrateInitialize(db, log)

	// todo init router: gin-gonic router
	// Logging to a file.
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	rep := repositories.NewRepository(db)
	service := services.NewService(rep)
	handlerCollection := handlers.NewHandler(service, log)

	// todo server: standart golang server
	srv := new(domains.Server)
	if err := srv.Run(cfg.HttpServer.Port, handlerCollection.InitRoutes()); err != nil {
		log.Debug("error occured while running http server: %s", err.Error())
	}
	fmt.Printf("server started in %s port", cfg.HttpServer.Port)
}
