package repositories

import (
	"crud/src/config"
	"crud/src/entities"
)

type GORMProductsRepository struct {}

func (repository GORMProductsRepository) GetAll() (*[]entities.Product, error) {
	var products []entities.Product

	result := config.Database.Select("id", "name", "price").Find(&products)

	return &products, result.Error
}

func (repository GORMProductsRepository) Create(product *entities.Product) (error) {
	result := config.Database.Create(product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
