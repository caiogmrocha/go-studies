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

func (repository GORMProductsRepository) GetOne(id uint) (*entities.Product, error) {
	var product entities.Product

	result := config.Database.Select("id", "name", "price").Where("id = ?", id).First(&product)

	return &product, result.Error
}

func (repository GORMProductsRepository) Create(product *entities.Product) (error) {
	result := config.Database.Create(product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository GORMProductsRepository) Update(product *entities.Product) (error) {
	result := config.Database.Model(product).Updates(entities.Product{
		Name: product.Name,
		Price: product.Price,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository GORMProductsRepository) Delete(product *entities.Product) (error) {
	result := config.Database.Delete(product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
