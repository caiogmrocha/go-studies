package repositories

import "crud/src/entities"

type ProductsRepository interface {
	GetAll() (*[]entities.Product, error)
	Create(product *entities.Product) (error)
}
