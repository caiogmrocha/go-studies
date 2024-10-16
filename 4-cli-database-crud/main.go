package main

import (
	_ "crud/src/config"
	"crud/src/controllers/cli"
	"fmt"
	"log"
	"os"
)

type Option int

const (
	ReadAll Option = iota
	ReadOne
	Create
	Update
	Delete
	Exit = -1
)

func main() {

	for {
		fmt.Println("Choose an option:")
		fmt.Println("0  - ReadAll")
		fmt.Println("1  - ReadOne")
		fmt.Println("2  - Create")
		fmt.Println("3  - Update")
		fmt.Println("4  - Delete")
		fmt.Println("-1 - Exit")

		var option Option

		fmt.Scanf("%d", &option)

		fmt.Printf("Selected option: %d\n\n", option)

		switch option {
			case ReadAll: {
				fmt.Println("Option not implemented")
			}

			case ReadOne: {
				fmt.Println("Option not implemented")
			}

			case Create: {
				productController := cli.ProductCliControllerFactory()

				err := productController.Create()

				if err != nil {
					log.Fatalf("Error while ProductCliController.Create(): %s", err.Error())
				}
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
				os.Exit(1)
			}
		}

		log.Print("App exited with status 0")
	}
}
