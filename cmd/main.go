package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
	"practice/domains"
	config "practice/internal/config"
	"practice/internal/logger"
	handlers "practice/pkg/handlers"
	"practice/pkg/repositories"
	"practice/pkg/services"
)

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

	// todo init router: gin-gonic router
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	rep := repositories.NewRepository()
	service := services.NewService(rep)
	handlerCollection := handlers.NewHandler(service, log)

	// todo server: standart golang server
	srv := new(domains.Server)
	if err := srv.Run("8081", handlerCollection.InitRoutes()); err != nil {
		log.Debug("error occured while running http server: %s", err.Error())
	}
}
