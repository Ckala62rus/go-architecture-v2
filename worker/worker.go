package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"practice/pkg/repositories"
	"practice/pkg/services"
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
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func GetServiceForTask() *services.Service {
	repo := repositories.NewRepository()
	return services.NewService(repo)
}

func work() {
	fmt.Println("i'm working!")
}

func work2() {
	service := GetServiceForTask()
	user := service.Users.GetUser()
	fmt.Printf("get user with name: %s \n", user.Name)
}
