package main

import (
	_ "crud/src/config"
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

func main() {
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
			case GetAll: {
				productController := cli.ProductCliControllerFactory()

				productController.GetAll()
			}

			case GetOne: {
				productController := cli.ProductCliControllerFactory()

				productController.GetOne()
			}

			case Create: {
				productController := cli.ProductCliControllerFactory()

				productController.Create()
			}

			case Update: {
				fmt.Println("Option not implemented")
			}

			case Delete: {
				fmt.Println("Option not implemented")
			}

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
