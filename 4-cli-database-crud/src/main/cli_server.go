package bootstrap

import (
	"crud/src/controllers/cli"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Option int

const (
	GetAll Option = iota
	GetOne
	Create
	Update
	Delete
	Exit = -1
)

func BootstrapCliServer() {
	productController := cli.ProductCliControllerFactory()

	for {
		fmt.Println("Choose an option:")
		fmt.Println("0  - GetAll")
		fmt.Println("1  - GetOne")
		fmt.Println("2  - Create")
		fmt.Println("3  - Update")
		fmt.Println("4  - Delete")
		fmt.Println("-1 - Exit")

		var option Option

		fmt.Scanf("%d", &option)

		fmt.Printf("Selected option: %d\n\n", option)

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
    cmd.Run()

		switch option {
			case GetAll: productController.GetAll()
			case GetOne: productController.GetOne()
			case Create: productController.Create()
			case Update: productController.Update()
			case Delete: productController.Delete()

			case Exit: {
				fmt.Println("App exited with status 0")
				os.Exit(0)
			}

			default: {
				log.Fatalf("Option not allowed: %d", option)
				log.Print("App exited with status 1")
				os.Exit(1)
			}
		}
	}
}
