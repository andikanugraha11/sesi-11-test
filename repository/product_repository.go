package repository

import "go-testing-11/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}


// Get Data to database