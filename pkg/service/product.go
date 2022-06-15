package service

import (
	"github.com/abdymazhit/go-test-task/model"
	"github.com/abdymazhit/go-test-task/pkg/repository"
)

type ProductService struct {
	encoderService *EncoderService
	repo           repository.Product
}

func NewProductService(encoderService *EncoderService, repo repository.Product) *ProductService {
	return &ProductService{encoderService, repo}
}

func (s *ProductService) GetNextId() (*int, error) {
	return s.repo.GetNextId()
}

func (s *ProductService) PutProduct(id int, product model.Product) error {
	productIdBytes, err := s.encoderService.encodeProductId(id)
	if err != nil {
		return nil
	}

	productBytes, err := s.encoderService.encodeProduct(product)
	if err != nil {
		return nil
	}

	return s.repo.PutProduct(productIdBytes, productBytes)
}

func (s *ProductService) GetProduct(id int) (*model.Product, error) {
	productIdBytes, err := s.encoderService.encodeProductId(id)
	if err != nil {
		return nil, err
	}

	productBytes, err := s.repo.GetProduct(productIdBytes)
	if err != nil {
		return nil, err
	}

	product, err := s.encoderService.decodeProduct(*productBytes)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetProducts() (*[]model.Product, error) {
	productsBytes, err := s.repo.GetProducts()
	if err != nil {
		return nil, err
	}

	var products []model.Product
	for _, productBytes := range *productsBytes {
		product, err := s.encoderService.decodeProduct(productBytes)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	return &products, nil
}

func (s *ProductService) DeleteProduct(id int) error {
	productIdBytes, err := s.encoderService.encodeProductId(id)
	if err != nil {
		return nil
	}
	return s.repo.DeleteProduct(productIdBytes)
}
