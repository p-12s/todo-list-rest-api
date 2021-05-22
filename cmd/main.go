package main

import (
	"github.com/p-12s/todo-list-rest-api"
	"github.com/p-12s/todo-list-rest-api/pkg/handler"
	"github.com/p-12s/todo-list-rest-api/pkg/repository"
	"github.com/p-12s/todo-list-rest-api/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
