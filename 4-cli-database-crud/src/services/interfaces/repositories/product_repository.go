package repositories

import "crud/src/entities"

type ProductsRepository interface {
	Create(product *entities.Product) (error)
}
