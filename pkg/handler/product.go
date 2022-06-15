package handler

import (
	"encoding/json"
	"github.com/abdymazhit/go-test-task/model"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) ProductAddHandler(ctx *fasthttp.RequestCtx) {
	var input model.AddProductRequest

	if err := json.Unmarshal(ctx.Request.Body(), &input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid json request")
		return
	}

	if input.Name == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "Product name cannot be empty")
		return
	}

	price, err := strconv.Atoi(input.Price)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Price must be a number")
		return
	}

	if price <= 0 {
		newErrorResponse(ctx, http.StatusBadRequest, "Price must be positive number")
		return
	}

	product := model.Product{
		Name:  input.Name,
		Price: price,
	}

	_, err = h.services.GetProductIdByName(product.Name)
	if err == nil { // when == nil it means the product with the same name already exists
		newErrorResponse(ctx, http.StatusBadRequest, "Product with the same name already exists")
		return
	}

	productId, err := h.services.GetNextId()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "Error getting next product id: "+err.Error())
		return
	}

	product.Id = *productId

	if err := h.services.PutProduct(product.Id, product); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "Error putting product to the bucket: "+err.Error())
		return
	}

	if err := h.services.PutProductIndex(product.Name, product.Id); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "Error putting product to the index bucket: "+err.Error())
		return
	}

	response, err := json.Marshal(model.GoodAddProductResponse{
		Status:    "success",
		ProductId: product.Id,
	})
	if err != nil {
		log.Println("Failed to send positive response: " + err.Error())
		return
	}

	_, _ = ctx.WriteString(string(response))
}

func (h *Handler) ProductEditHandler(ctx *fasthttp.RequestCtx) {
	var input model.EditProductRequest

	if err := json.Unmarshal(ctx.Request.Body(), &input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid json request")
		return
	}

	if input.Name == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "Product name cannot be empty")
		return
	}

	price, err := strconv.Atoi(input.Price)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Price must be a number")
		return
	}

	if price <= 0 {
		newErrorResponse(ctx, http.StatusBadRequest, "Price must be positive number")
		return
	}

	product := model.Product{
		Id:    input.Id,
		Name:  input.Name,
		Price: price,
	}

	oldProduct, err := h.services.GetProduct(product.Id)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Product with this id does not exist")
		return
	}

	_, err = h.services.GetProductIdByName(product.Name)
	if err == nil { // when == nil it means the product with the same name already exists
		newErrorResponse(ctx, http.StatusBadRequest, "Product with the same name already exists")
		return
	}

	if err := h.services.DeleteProductIndex(oldProduct.Name); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "Error deleting old product from the index bucket: "+err.Error())
		return
	}

	if err := h.services.PutProduct(product.Id, product); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "Error putting product to the bucket: "+err.Error())
		return
	}

	if err := h.services.PutProductIndex(product.Name, product.Id); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "Error putting product to the index bucket: "+err.Error())
		return
	}

	response, err := json.Marshal(model.GoodResponse{
		Status: "success",
	})
	if err != nil {
		log.Println("Failed to send positive response: " + err.Error())
		return
	}

	_, _ = ctx.WriteString(string(response))
}

func (h *Handler) ProductDeleteHandler(ctx *fasthttp.RequestCtx) {
	var input model.DeleteProductRequest

	if err := json.Unmarshal(ctx.Request.Body(), &input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid json request")
		return
	}

	product, err := h.services.GetProduct(input.Id)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Product with this id does not exist")
		return
	}

	if err := h.services.DeleteProduct(product.Id); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "Error deleting product from the bucket: "+err.Error())
		return
	}

	if err := h.services.DeleteProductIndex(product.Name); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "Error deleting product from the index bucket: "+err.Error())
		return
	}

	response, err := json.Marshal(model.GoodResponse{
		Status: "success",
	})
	if err != nil {
		log.Println("Failed to send positive response: " + err.Error())
		return
	}

	_, _ = ctx.WriteString(string(response))
}

func (h *Handler) ProductSearchHandler(ctx *fasthttp.RequestCtx) {
	var input model.SearchProductRequest

	if err := json.Unmarshal(ctx.Request.Body(), &input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid json request")
		return
	}

	productId, err := h.services.GetProductIdByName(input.SearchName)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Product with this name does not exist")
		return
	}

	product, err := h.services.GetProduct(*productId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "Product with this id does not exist")
		return
	}

	response, err := json.Marshal(model.GoodProductResponse{
		Status:  "success",
		Product: *product,
	})
	if err != nil {
		log.Println("Failed to send positive response: " + err.Error())
		return
	}

	_, _ = ctx.WriteString(string(response))
}
