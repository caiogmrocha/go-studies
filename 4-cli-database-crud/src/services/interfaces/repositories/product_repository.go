package repositories

import "crud/src/entities"

type ProductRepository interface {
	Create(product *entities.Product) (error)
}
