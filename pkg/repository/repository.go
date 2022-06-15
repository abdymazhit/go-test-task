package repository

import (
	"github.com/boltdb/bolt"
)

type Product interface {
	GetNextId() (*int, error)
	PutProduct(id, product []byte) error
	GetProduct(id []byte) []byte
	DeleteProduct(id []byte) error
}

type ProductIndex interface {
	PutProductIndex(name, id []byte) error
	GetProductIdByName(name []byte) []byte
	DeleteProductIndex(name []byte) error
}

type Repository struct {
	Product
	ProductIndex
}

func NewRepository(productsBucket, productIndexBucket *bolt.Bucket) *Repository {
	return &Repository{
		Product:      NewProductRepo(productsBucket),
		ProductIndex: NewProductIndex(productIndexBucket),
	}
}
