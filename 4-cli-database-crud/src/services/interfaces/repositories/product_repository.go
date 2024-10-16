package repositories

import "crud/src/entities"

type ProductsRepository interface {
	GetAll() (*[]entities.Product, error)
	GetOne(id int) (*entities.Product, error)
	Create(product *entities.Product) (error)
	Update(product *entities.Product) (error)
	Delete(product *entities.Product) (error)
}
