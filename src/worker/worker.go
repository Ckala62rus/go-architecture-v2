package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"practice/internal/config"
	"practice/pkg/repositories"
	"practice/pkg/services"

	"github.com/robfig/cron/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("hello cron task!")

	c := cron.New(
		cron.WithParser(
			cron.NewParser(
				cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))

	_, err := c.AddFunc("*/10 * * * * *", work)
	if err != nil {
		fmt.Println(err)
	}

	c.AddFunc("*/2 * * * * *", work2)

	go c.Start()

	// for run forever
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func GetDB() *gorm.DB {
	projectDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cfg := config.MustLoad(filepath.Join(projectDir, "config", "config.yml"))
	//fmt.Printf("%+v", cfg) // todo need delete!

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
	//repo := repositories.NewRepository(db)
	//return services.NewService(repo)
	return db
}

func work() {
	fmt.Println("i'm working!")
}

func work2() {
	db := GetDB()
	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()

	repo := repositories.NewRepository(db)
	service := services.NewService(repo)

	user := service.Users.GetAllUsers()
	fmt.Printf("get user with name: %v \n", user)
}
