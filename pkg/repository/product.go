package repository

import (
	"github.com/boltdb/bolt"
)

type ProductRepo struct {
	db *bolt.DB
}

func NewProductRepo(db *bolt.DB) *ProductRepo {
	return &ProductRepo{db}
}

func (r *ProductRepo) GetNextId() (*int, error) {
	var productId int

	if err := r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("product"))

		id, err := b.NextSequence()
		if err != nil {
			return err
		}
		productId = int(id)

		return nil
	}); err != nil {
		return nil, err
	}

	return &productId, nil
}

func (r *ProductRepo) PutProduct(idBytes, productBytes []byte) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("product"))
		return b.Put(idBytes, productBytes)
	})
}

func (r *ProductRepo) GetProduct(idBytes []byte) (*[]byte, error) {
	var result []byte

	if err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("product"))
		result = b.Get(idBytes)
		return nil
	}); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ProductRepo) GetProducts() (*[][]byte, error) {
	var values [][]byte

	if err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("product"))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			values = append(values, v)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &values, nil
}

func (r *ProductRepo) DeleteProduct(idBytes []byte) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("product"))
		return b.Delete(idBytes)
	})
}
