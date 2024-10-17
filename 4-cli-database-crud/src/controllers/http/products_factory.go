package controllers

import (
	repositories "crud/src/repositories/gorm"
	"crud/src/services"
)

func ProductHttpControllerFactory() (ProductsHttpController) {
	ProductController := ProductsHttpController{
		ProductsService: services.ProductsService{
			ProductsRepository: repositories.GORMProductsRepository{},
		},
	}

	return ProductController
}
