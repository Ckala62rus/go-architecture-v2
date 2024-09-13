package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"practice/pkg/utils"
	"syscall"
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

	cfg := utils.MainConfig

	MainLogger = SetupNewLogger(cfg.Env, filepath.Join(projectDir, "logs", "logs.log"))
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
	case envDev:
		log = slog.New(slog.NewJSONHandler(f, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	//log = slog.New(slog.NewJSONHandler(f, &slog.HandlerOptions{Level: slog.LevelDebug}))

	return log
}
