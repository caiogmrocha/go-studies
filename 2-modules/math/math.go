package math

import "errors"

func Fatorial(n int) (int, error) {
	if (n < 0) {
		return 0, errors.New("n is an signed number")
	} else if (n < 2) {
		return 1, nil
	}

	f, error := Fatorial(n - 1)

	if (error != nil) {
		return 0, error
	}

	return (n * f), nil
}

func Sum(numbers *[]float64) float64 {
  sum := 0.

  for _, n := range *numbers {
    sum += n
  }

  return sum
}

func SquareArray(numbers *[]float64) {
	for k, v := range *numbers {
		(*numbers)[k] = v*v
	}
}

func PrimesLessThan(l int) ([]int, error) {
	if (l < 2) {
		return []int{}, errors.New("l is less than 2")
	}

	primes := []int{}

	loop: for i := 2; i < l; i++ {
		for j := 2; j < i; j++ {
			if i % j == 0 {
				continue loop
			}
		}

		primes = append(primes, i)
	}

	return primes, nil
}
