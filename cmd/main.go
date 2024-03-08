package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"practice/domains"
	config "practice/internal/config"
	handlers "practice/pkg/handlers"
	"practice/pkg/repositories"
	"practice/pkg/services"
	"syscall"
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
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	f := &lumberjack.Logger{
		Filename:   "logs/logs.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	go func() {
		for {
			<-c
			f.Rotate()
		}
	}()

	switch env {
	case envLocal:
		//log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		log = slog.New(slog.NewJSONHandler(f, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(f, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	//log = slog.New(slog.NewJSONHandler(f, &slog.HandlerOptions{Level: slog.LevelDebug}))

	return log
}
