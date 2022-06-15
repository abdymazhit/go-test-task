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

	r.GET("/q/product-search-by-name", h.ProductSearchHandler)

	r.POST("/cmd/add-product", h.ProductAddHandler)
	r.PUT("/cmd/edit-product", h.ProductEditHandler)
	r.DELETE("/cmd/delete-product", h.ProductDeleteHandler)

	return r
}
