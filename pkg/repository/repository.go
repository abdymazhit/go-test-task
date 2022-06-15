package repository

import (
	"github.com/boltdb/bolt"
)

type Product interface {
	GetNextId() (*int, error)
	PutProduct(id, product []byte) error
	GetProduct(idBytes []byte) (*[]byte, error)
	DeleteProduct(id []byte) error
}

type ProductIndex interface {
	PutProductIndex(name, id []byte) error
	GetProductIdByName(nameBytes []byte) (*[]byte, error)
	DeleteProductIndex(name []byte) error
}

type Repository struct {
	Product
	ProductIndex
}

func NewRepository(db *bolt.DB) *Repository {
	return &Repository{
		Product:      NewProductRepo(db),
		ProductIndex: NewProductIndex(db),
	}
}
