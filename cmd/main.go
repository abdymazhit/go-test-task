package main

import (
	"github.com/abdymazhit/go-test-task/pkg/handler"
	"github.com/abdymazhit/go-test-task/pkg/repository"
	"github.com/abdymazhit/go-test-task/pkg/service"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	r := handlers.InitRoutes()
	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
