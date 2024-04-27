package main

import (
	"first-rest-api/pkg/handler"
	"first-rest-api/pkg/service"
	"log"

	"github.com/spf13/viper"
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http sever: %s", err.Error())
	}
}

func initConfig() error {
	 viper.AddConfigPath("configs")
	 viper.SetConfigName("config")
	 return viper.ReadInConfig()
}