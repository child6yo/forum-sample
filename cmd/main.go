package main

import (
    "github.com/child6yo/forum-sample"

    "github.com/child6yo/forum-sample/pkg/repository"
    "github.com/child6yo/forum-sample/pkg/service"
    "github.com/child6yo/forum-sample/pkg/handler"
)

func main() {
    srv := new(forum.Server)
    repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
    srv.Run("8000", handlers.InitRoutes())
}