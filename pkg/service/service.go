package service

import (
	"github.com/abdymazhit/go-test-task/model"
	"github.com/abdymazhit/go-test-task/pkg/repository"
)

type Encoder interface {
	encodeProductId(id int) ([]byte, error)
	decodeProductId(productIdBytes []byte) (*int, error)
	encodeProductName(name string) ([]byte, error)
	decodeProductName(productNameBytes []byte) (*string, error)
	encodeProduct(product model.Product) ([]byte, error)
	decodeProduct(productBytes []byte) (*model.Product, error)
}

type Product interface {
	GetNextId() (*int, error)
	PutProduct(id int, product model.Product) error
	GetProduct(id int) (*model.Product, error)
	DeleteProduct(id int) error
}

type ProductIndex interface {
	PutProductIndex(name string, id int) error
	GetProductIdByName(name string) (*int, error)
	DeleteProductIndex(name string) error
}

type Service struct {
	Product
	ProductIndex
}

func NewService(repos *repository.Repository) *Service {
	encoderService := NewEncoderService()

	return &Service{
		Product:      NewProductService(encoderService, repos.Product),
		ProductIndex: NewProductIndexService(encoderService, repos.ProductIndex),
	}
}
