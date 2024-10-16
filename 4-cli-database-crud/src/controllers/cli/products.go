package cli

import (
	"bufio"
	"crud/src/services"
	"fmt"
	"log"
	"os"
	"strings"
)

type ProductsCliController struct {
	ProductsService services.ProductsService
}

func (controller *ProductsCliController) GetAll() {
	products, err := controller.ProductsService.GetAll()

	if err != nil {
		log.Fatalf("Error while ProductCliController.GetAll(): %s", err.Error())
	}

	fmt.Print("\nProdutos:\n\n")

	for _, product := range *products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n\n", product.ID, product.Name, product.Price)
	}
}

func (controller *ProductsCliController) Create() {
	fmt.Print("Informe os dados do produto:\n\n")

	dto := new(services.CreateProductDto)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')

	dto.Name = strings.Replace(name, "\n", "", 1)

	fmt.Print("Price: ")
	fmt.Scanf("%f", &dto.Price)

	err := controller.ProductsService.Create(dto)

	if err != nil {
		log.Fatalf("Error while ProductCliController.Create(): %s", err.Error())
	}
}
