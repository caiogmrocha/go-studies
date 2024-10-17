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
		log.Printf("Error while ProductsHttpController.GetAll(): %s", err.Error())

		c.JSON(500, gin.H{
			"message": "Internal server error",
		})

		return
	}

	c.JSON(200, products)
}

func (controller *ProductsHttpController) GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		log.Printf("Error while ProductsHttpController.GetAll(): %s", err.Error())

		c.JSON(500, gin.H{
			"message": "Internal server error",
		})

		return
	}

	products, err := controller.ProductsService.GetOne(id)

	if err != nil {
		log.Printf("Error while ProductsHttpController.GetAll(): %s", err.Error())

		c.JSON(500, gin.H{
			"message": "Internal server error",
		})

		return
	}

	c.JSON(200, products)
}

func (controller *ProductsHttpController) Create(c *gin.Context) {
	parsedBody := make(map[string]interface{})

	err := c.BindJSON(&parsedBody)

	if err != nil {
		log.Printf("Error while ProductsHttpController.Create(): %s", err.Error())

		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})

		return
	}

	name := parsedBody["name"].(string)
	price := parsedBody["price"].(float64)

	dto := services.CreateProductDto{
		Name: name,
		Price: price,
	}

	err = controller.ProductsService.Create(&dto)

	if err != nil {
		log.Printf("Error while ProductsHttpController.Create(): %s", err.Error())

		c.JSON(500, gin.H{
			"message": "Internal server error",
		})

		return
	}

	c.JSON(201, gin.H{
		"message": "Product created successfully",
	})
}
