package repository

import "github.com/boltdb/bolt"

type ProductIndexRepo struct {
	db *bolt.DB
}

func NewProductIndex(db *bolt.DB) *ProductIndexRepo {
	return &ProductIndexRepo{db}
}

func (r *ProductIndexRepo) PutProductIndex(nameBytes, idBytes []byte) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("productIndex"))
		return b.Put(nameBytes, idBytes)
	})
}

func (r *ProductIndexRepo) GetProductIdByName(nameBytes []byte) (*[]byte, error) {
	var result []byte

	if err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("productIndex"))
		result = b.Get(nameBytes)
		return nil
	}); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ProductIndexRepo) DeleteProductIndex(nameBytes []byte) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("productIndex"))
		return b.Delete(nameBytes)
	})
}
