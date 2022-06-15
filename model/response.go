package model

type GoodResponse struct {
	Status string `json:"status"`
}

type GoodAddProductResponse struct {
	Status    string `json:"status"`
	ProductId int    `json:"productId"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
