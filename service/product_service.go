package service

import (
	"errors"
	"go-testing-11/entity"
	"go-testing-11/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (s ProductService) GetProduct(id string) (*entity.Product, error) {
	product := s.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}