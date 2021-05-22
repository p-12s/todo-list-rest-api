package main

import (
	"github.com/p-12s/todo-list-rest-api"
	"github.com/p-12s/todo-list-rest-api/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
