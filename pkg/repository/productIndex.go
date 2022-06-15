package repository

import "github.com/boltdb/bolt"

type ProductIndexRepo struct {
	bucket *bolt.Bucket
}

func NewProductIndex(bucket *bolt.Bucket) *ProductIndexRepo {
	return &ProductIndexRepo{bucket}
}

func (r *ProductIndexRepo) PutProductIndex(nameBytes, idBytes []byte) error {
	return r.bucket.Put(nameBytes, idBytes)
}

func (r *ProductIndexRepo) GetProductIdByName(nameBytes []byte) []byte {
	return r.bucket.Get(nameBytes)
}

func (r *ProductIndexRepo) DeleteProductIndex(nameBytes []byte) error {
	return r.bucket.Delete(nameBytes)
}
