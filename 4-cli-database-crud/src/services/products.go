package services

import (
	"crud/src/entities"
	"crud/src/services/interfaces/repositories"
)

type ProductsService struct {
	ProductsRepository repositories.ProductsRepository
}

func (service *ProductsService) GetAll() (*[]entities.Product, error) {
	products, err := service.ProductsRepository.GetAll()

	return products, err
}

func (service *ProductsService) GetOne(id uint) (*entities.Product, error) {
	product, err := service.ProductsRepository.GetOne(id)

	return product, err
}

type CreateProductDto struct {
	Name string
	Price float64
}

func (service *ProductsService) Create(dto *CreateProductDto) (error) {
	product := &entities.Product{
		Name: dto.Name,
		Price: dto.Price,
	}

	err := service.ProductsRepository.Create(product)

	return err
}

func (service *ProductsService) Update(product *entities.Product) (error) {
	err := service.ProductsRepository.Update(product)

	return err
}

func (service *ProductsService) Delete(product *entities.Product) (error) {
	err := service.ProductsRepository.Delete(product)

	return err
}
