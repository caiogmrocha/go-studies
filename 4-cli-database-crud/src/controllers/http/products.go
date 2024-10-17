package controllers

import (
	"crud/src/services"
	"log"
	"strconv"

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

func (controller *ProductsHttpController) GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		log.Fatalf("Error while ProductsHttpController.GetAll(): %s", err.Error())
	}

	products, err := controller.ProductsService.GetOne(id)

	if err != nil {
		log.Fatalf("Error while ProductsHttpController.GetAll(): %s", err.Error())
	}

	c.JSON(200, products)
}
