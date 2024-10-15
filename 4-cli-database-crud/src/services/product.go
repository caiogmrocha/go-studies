package services

import (
	"crud/src/entities"
	"crud/src/services/interfaces/repositories"
)

type ProductService struct {
	productRepository repositories.ProductRepository
}

func (service *ProductService) Create(product *entities.Product) (error) {
	err := service.productRepository.Create(product)

	return err
}
