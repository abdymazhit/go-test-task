package repository

import (
	"github.com/boltdb/bolt"
)

type ProductRepo struct {
	bucket *bolt.Bucket
}

func NewProductRepo(bucket *bolt.Bucket) *ProductRepo {
	return &ProductRepo{bucket}
}

func (r *ProductRepo) GetNextId() (*int, error) {
	id, err := r.bucket.NextSequence()
	if err != nil {
		return nil, err
	}

	productId := int(id)
	return &productId, nil
}

func (r *ProductRepo) PutProduct(idBytes, productBytes []byte) error {
	return r.bucket.Put(idBytes, productBytes)
}

func (r *ProductRepo) GetProduct(idBytes []byte) []byte {
	return r.bucket.Get(idBytes)
}

func (r *ProductRepo) DeleteProduct(idBytes []byte) error {
	return r.bucket.Delete(idBytes)
}
