package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("can't get square root of negative number")
	}

	return math.Sqrt(number), nil
}

func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by zero")
	}

	return dividend / divisor, nil
}

func main() {
	root, err := squareRoot(4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%0.3f \n", root)
}
