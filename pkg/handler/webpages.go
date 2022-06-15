package handler

import (
	"fmt"
	"github.com/abdymazhit/go-test-task/templates"
	"github.com/valyala/fasthttp"
	"strconv"
)

func (h *Handler) ProductsPageHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")

	products, err := h.services.GetProducts()
	if err != nil {
		p := &templates.ErrPage{
			Error: "Error getting all products: " + err.Error(),
		}
		p.WriteBody(ctx)
		return
	}

	p := &templates.ProductsPage{
		Products: *products,
	}
	p.WriteBody(ctx)
}

func (h *Handler) ProductAddPageHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	p := &templates.ProductAddPage{}
	p.WriteBody(ctx)
}

func (h *Handler) ProductEditPageHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")

	id := ctx.UserValue("productId")
	productId, err := strconv.Atoi(fmt.Sprintf("%s", id))
	if err != nil {
		p := &templates.ErrPage{
			Error: "Error getting product by id: " + err.Error(),
		}
		p.WriteBody(ctx)
		return
	}

	product, err := h.services.GetProduct(productId)
	if err != nil {
		p := &templates.ErrPage{
			Error: "Error getting product by id: " + err.Error(),
		}
		p.WriteBody(ctx)
		return
	}

	p := &templates.ProductEditPage{
		Name:  product.Name,
		Price: product.Price,
	}
	p.WriteBody(ctx)
}
