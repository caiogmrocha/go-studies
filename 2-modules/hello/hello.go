package main

import (
	"fmt"

	"example.com/math"
)
func main() {
	// Fatorial test
	f, error := math.Fatorial(5)

	if (error != nil) {
		fmt.Print("Fatorial de 5: ")
    fmt.Println(error)
	} else {
		fmt.Println(f)
	}

	numbers := []float64{1,2,3,4}

	// Sum test
	sum := math.Sum(&numbers)

	fmt.Println(sum)

	// SquareArray test
	math.SquareArray(&numbers)

	fmt.Print("[ ")

	for i := 0; i < len(numbers); i++ {
		if (i != len(numbers) - 1) {
			fmt.Printf("%v, ", numbers[i])
		} else {
			fmt.Printf("%v", numbers[i])
		}
	}

	fmt.Println(" ]")
}
