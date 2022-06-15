package service

import "github.com/abdymazhit/go-test-task/pkg/repository"

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
