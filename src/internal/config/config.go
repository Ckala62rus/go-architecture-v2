package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env            string `yaml:"env" env:"ENV" env-default:"development" env-required:"true"`
	ConfigFile     string `yaml:"config_file" env:"config_file" env-default:"local" env-required:"true"`
	HttpServer     `yaml:"http_server"`
	DatabaseConfig `yaml:"database"`
}

type HttpServer struct {
	Address     string        `yaml:"address" evn-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	Port        string        `yaml:"port" env-default:"8081"`
}

type DatabaseConfig struct {
	Host     string `yaml:"postgres_host"`
	Port     string `yaml:"postgres_port"`
	User     string `yaml:"postgres_user"`
	Password string `yaml:"postgres_password"`
	Db       string `yaml:"postgres_db"`
}

func MustLoad(configPath string) *Config {
	//configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
