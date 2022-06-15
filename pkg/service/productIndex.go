package service

import (
	"github.com/abdymazhit/go-test-task/pkg/repository"
)

type ProductIndexService struct {
	encoderService *EncoderService
	repo           repository.ProductIndex
}

func NewProductIndexService(encoderService *EncoderService, repo repository.ProductIndex) *ProductIndexService {
	return &ProductIndexService{encoderService, repo}
}

func (s *ProductIndexService) PutProductIndex(name string, id int) error {
	productNameBytes, err := s.encoderService.encodeProductName(name)
	if err != nil {
		return nil
	}

	productIdBytes, err := s.encoderService.encodeProductId(id)
	if err != nil {
		return nil
	}

	return s.repo.PutProductIndex(productNameBytes, productIdBytes)
}

func (s *ProductIndexService) GetProductIdByName(name string) (*int, error) {
	productNameBytes, err := s.encoderService.encodeProductName(name)
	if err != nil {
		return nil, err
	}

	productIdBytes := s.repo.GetProductIdByName(productNameBytes)
	productId, err := s.encoderService.decodeProductId(productIdBytes)
	if err != nil {
		return nil, err
	}

	return productId, nil
}

func (s *ProductIndexService) DeleteProductIndex(name string) error {
	productNameBytes, err := s.encoderService.encodeProductName(name)
	if err != nil {
		return nil
	}
	return s.repo.DeleteProductIndex(productNameBytes)
}
