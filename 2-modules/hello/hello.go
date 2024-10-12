package main

import (
	"fmt"

	"example.com/math"
)
func main() {
	f, error := math.Fatorial(5)

	if (error != nil) {
		fmt.Print("Fatorial de 5: ")
    fmt.Println(error)
	} else {
		fmt.Println(f)
	}

	numbers := []float64{1,2,3,4}

	sum := math.Sum(&numbers)

	fmt.Println(sum)
}
