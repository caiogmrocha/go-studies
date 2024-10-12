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

	fmt.Println(numbers)

	// PrimesLessThan test

	primes, error := math.PrimesLessThan(20)

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(primes)
}
