package main

import (
	"fmt"

	"example.com/math"
)
func main() {
	f, error := math.Fatorial(5)

	if (error != nil) {
		fmt.Println(error)
	} else {
		fmt.Println(f)
	}
}
