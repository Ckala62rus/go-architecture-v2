package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	log2 "log"
	"log/slog"
	"os"
	"path/filepath"
	"practice/domains"
	config "practice/internal/config"
	handlers "practice/pkg/handlers"
	"practice/pkg/repositories"
	"practice/pkg/services"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	projectDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// todo init config: cleanenv
	cfg := config.MustLoad(filepath.Join(projectDir, "config", "local.yml"))
	//fmt.Printf("%+v", cfg) // todo need delete!

	// todo init logger: slog
	log := setupLogger(cfg.Env)
	log.Debug("Hello slog logger!")

	// todo init storage: gorm
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.User,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.Db,
		cfg.DatabaseConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Debug()

	//var categories = []domains.Category{}
	////db.Debug().First(&category, 7)
	//db.Debug().Find(&categories)
	//fmt.Printf("%+v", categories)

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// todo init router: gin-gonic router
	rep := repositories.NewRepository()
	service := services.NewService(rep)
	handlerCollection := handlers.NewHandler(service)

	// todo server: standart golang server
	srv := new(domains.Server)
	if err := srv.Run("8081", handlerCollection.InitRoutes()); err != nil {
		log.Debug("error occured while running http server: %s", err.Error())
	}

	//user := service.Users.GetUser()
	//newUser := service.Users.CreateUser("Evgeniy", 34)
	//
	//fmt.Printf("user with name %s \n", user.Name)
	//fmt.Printf("user with name %s \n", newUser.Name)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	f, err := os.OpenFile("logging", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log2.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	switch env {
	//case envLocal:
	//	log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envLocal:
		log = slog.New(slog.NewJSONHandler(f, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
