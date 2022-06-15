package model

type GoodResponse struct {
	Status string `json:"status"`
}

type GoodAddProductResponse struct {
	Status    string `json:"status"`
	ProductId int    `json:"productId"`
}

type GoodProductResponse struct {
	Status  string  `json:"status"`
	Product Product `json:"product"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
