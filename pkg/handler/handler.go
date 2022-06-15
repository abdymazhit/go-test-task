package handler

import (
	"github.com/abdymazhit/go-test-task/pkg/service"
	"github.com/fasthttp/router"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes() *router.Router {
	r := router.New()

	return r
}
