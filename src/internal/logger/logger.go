package logger

import (
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

var (
	MainLogger *slog.Logger
)

func init() {
	projectDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Используем ENV переменную или значение по умолчанию
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	MainLogger = SetupNewLogger(env, filepath.Join(projectDir, "logs", "logs.log"))
	MainLogger.Info("***************** LOGGER INITIALIZED RUN *****************")
}

func SetupNewLogger(env string, path string) *slog.Logger {
	var log *slog.Logger

	f := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    5, // megabytes
		MaxBackups: 8,
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
	case envDev, "development":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd, "production":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		// Fallback to dev mode for unknown environments
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	//log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})) // log to console
	//log = slog.New(slog.NewJSONHandler(f, &slog.HandlerOptions{Level: slog.LevelDebug}))

	return log
}
