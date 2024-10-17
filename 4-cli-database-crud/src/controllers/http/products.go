package controllers

import (
	"crud/src/services"
	"log"

	"github.com/gin-gonic/gin"
)

type ProductsHttpController struct {
	ProductsService services.ProductsService
}

func (controller *ProductsHttpController) GetAll(c *gin.Context) {
	products, err := controller.ProductsService.GetAll()

	if err != nil {
		log.Fatalf("Error while ProductsHttpController.GetAll(): %s", err.Error())
	}

	c.JSON(200, products)
}
