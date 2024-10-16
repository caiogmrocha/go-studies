package cli

import (
	repositories "crud/src/repositories/gorm"
	"crud/src/services"
)

func ProductCliControllerFactory() (ProductsCliController) {
	ProductController := ProductsCliController{
		ProductsService: services.ProductsService{
			ProductsRepository: repositories.GORMProductsRepository{},
		},
	}

	return ProductController
}
