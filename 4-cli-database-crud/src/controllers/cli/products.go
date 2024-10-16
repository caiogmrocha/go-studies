package cli

import (
	"bufio"
	"crud/src/services"
	"fmt"
	"os"
	"strings"
)

type ProductsCliController struct {
	ProductsService services.ProductsService
}

func (controller *ProductsCliController) Create() (error) {
	fmt.Print("Informe os dados do produto:\n\n")

	dto := new(services.CreateProductDto)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')

	dto.Name = strings.Replace(name, "\n", "", 1)

	fmt.Print("Price: ")
	fmt.Scanf("%f", &dto.Price)

	err := controller.ProductsService.Create(dto)

	return err
}
