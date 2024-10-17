package controllers

import (
	"crud/src/entities"
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

	products, err := controller.ProductsService.GetOne(uint(id))

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

func (controller *ProductsHttpController) Update(c *gin.Context) {
	parsedBody := make(map[string]interface{})

	err := c.BindJSON(&parsedBody)

	if err != nil {
		log.Printf("Error while ProductsHttpController.Update(): %s", err.Error())

		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})

		return
	}

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)

	if err != nil {
		log.Printf("Error while ProductsHttpController.Update(): %s", err.Error())

		c.JSON(500, gin.H{
			"message": "Internal server error",
		})

		return
	}

	name := parsedBody["name"].(string)
	price := parsedBody["price"].(float64)

	dto := entities.Product{
		ID: uint(id),
		Name: name,
		Price: price,
	}

	err = controller.ProductsService.Update(&dto)

	if err != nil {
		log.Printf("Error while ProductsHttpController.Update(): %s", err.Error())

		c.JSON(500, gin.H{
			"message": "Internal server error",
		})

		return
	}

	c.JSON(201, gin.H{
		"message": "Product updated successfully",
	})
}

func (controller *ProductsHttpController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)

	if err != nil {
		log.Printf("Error while ProductsHttpController.Delete(): %s", err.Error())

		c.JSON(500, gin.H{
			"message": "Internal server error",
		})

		return
	}

	product, err := controller.ProductsService.GetOne(uint(id))

	if err != nil {
		log.Printf("Error while ProductsHttpController.Delete(): %s", err.Error())

		c.JSON(404, gin.H{
			"message": "Product not found",
		})

		return
	}

	err = controller.ProductsService.Delete(product)

	if err != nil {
		log.Printf("Error while ProductsHttpController.Delete(): %s", err.Error())

		c.JSON(500, gin.H{
			"message": "Internal server error",
		})

		return
	}

	c.JSON(201, gin.H{
		"message": "Product deleted successfully",
	})
}
