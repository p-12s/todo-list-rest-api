package main

import (
	"github.com/p-12s/todo-list-rest-api"
	"github.com/p-12s/todo-list-rest-api/pkg/handler"
	"github.com/p-12s/todo-list-rest-api/pkg/repository"
	"github.com/p-12s/todo-list-rest-api/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init configs: %s\n", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}