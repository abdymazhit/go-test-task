package model

type AddProductRequest struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

type EditProductRequest struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type DeleteProductRequest struct {
	Id int `json:"id"`
}

type SearchProductRequest struct {
	SearchName string `json:"searchName"`
}
