package bootstrap

import (
	controllers "crud/src/controllers/http"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func BootstrapHttpServer() {
	r := gin.Default()

	productsHttpController := controllers.ProductHttpControllerFactory()

	r.GET("/api/products", productsHttpController.GetAll)
	r.GET("/api/products/:id", productsHttpController.GetOne)
	r.POST("/api/products", productsHttpController.Create)

	err := r.Run(":8080")

	if err != nil {
		log.Fatalf("Failed to up server at :8080")
		os.Exit(1)
	}

	log.Println("Server is up at port :8080")
}
