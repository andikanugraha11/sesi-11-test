package service_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-testing-11/entity"
	"go-testing-11/repository"
	"go-testing-11/service"
	"testing"
)


var productRepository = &repository.ProductRepositoryMock{Mock:mock.Mock{}}
var productService = service.ProductService{Repository:productRepository}


func TestProductService_GetProductNotFound(t *testing.T) {
	// id 1 --> data == nil

	//Mocking data
	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)

	assert.Equal(t, "product not found", err.Error(), "error should product not found")
}

func TestProductService_GetProduct(t *testing.T) {
	// id 2 --> data == product
	productMock := entity.Product{
		Id:   "2",
		Name: "Product 2",
	}

	// Mocking data
	productRepository.Mock.On("FindById", "2").Return(productMock)

	product, err := productService.GetProduct("2")

	assert.Nil(t, err)

	assert.NotNil(t, product)

	assert.Equal(t, &productMock, product, "result must be equals")
}