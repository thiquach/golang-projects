package main

import (
	"fmt"
	"math"
)

// Finds square root of x
// starting with z = 1
// adjust z using the equation z -= (z*z - x) / (2*z)
// compare z*z with x and stop once the value has stopped changing
func Sqrt(x float64) float64 {
	var z float64
	var a float64

	z = float64(1)

	for math.Abs(float64(x-z*z)) > 0.001 {
		a = float64(z*z-x) / float64(2*z)
		z -= a
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
