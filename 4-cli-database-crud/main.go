package main

import (
	_ "crud/src/config"
	bootstrap "crud/src/main"
)

func main() {
	bootstrap.BootstrapHttpServer()
}
