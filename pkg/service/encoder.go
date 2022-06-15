package service

import (
	"bytes"
	"encoding/gob"
	"github.com/abdymazhit/go-test-task/model"
)

type EncoderService struct {
}

func NewEncoderService() *EncoderService {
	return &EncoderService{}
}

func (s *EncoderService) encodeProductId(id int) ([]byte, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)

	if err := enc.Encode(id); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (s *EncoderService) decodeProductId(productIdBytes []byte) (*int, error) {
	buff := bytes.NewBuffer(productIdBytes)
	dec := gob.NewDecoder(buff)

	var productId int
	if err := dec.Decode(&productId); err != nil {
		return nil, err
	}

	return &productId, nil
}

func (s *EncoderService) encodeProductName(name string) ([]byte, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)

	if err := enc.Encode(name); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (s *EncoderService) decodeProductName(productNameBytes []byte) (*string, error) {
	buff := bytes.NewBuffer(productNameBytes)
	dec := gob.NewDecoder(buff)

	var productName string
	if err := dec.Decode(&productName); err != nil {
		return nil, err
	}

	return &productName, nil
}

func (s *EncoderService) encodeProduct(product model.Product) ([]byte, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)

	if err := enc.Encode(product); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (s *EncoderService) decodeProduct(productBytes []byte) (*model.Product, error) {
	buff := bytes.NewBuffer(productBytes)
	dec := gob.NewDecoder(buff)

	var product model.Product
	if err := dec.Decode(&product); err != nil {
		return nil, err
	}

	return &product, nil
}
