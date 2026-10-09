package main

import (
	"errors"
	"fmt"
	"math"
)

// Finds square root of x using Newton's method
func Sqrt(x float64) (float64, error) {
	var z float64
	var a float64

	if x == 0 || x < 0 {
		return x, errors.New("invalid value")
	}

	z = float64(1)

	for math.Abs(float64(x-z*z)) > 0.001 {
		a = float64(z*z-x) / float64(2*z)
		z -= a
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
