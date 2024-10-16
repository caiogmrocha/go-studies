package services

import (
	"crud/src/entities"
	"crud/src/services/interfaces/repositories"
)

type ProductsService struct {
	ProductsRepository repositories.ProductsRepository
}

type CreateProductDto struct {
	Name string
	Price float64
}

func (service *ProductsService) GetAll() (*[]entities.Product, error) {
	products, err := service.ProductsRepository.GetAll()

	return products, err
}


func (service *ProductsService) Create(dto *CreateProductDto) (error) {
	product := &entities.Product{
		Name: dto.Name,
		Price: dto.Price,
	}

	err := service.ProductsRepository.Create(product)

	return err
}
