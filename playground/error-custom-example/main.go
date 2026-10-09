package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Finds square root of x using Newton's method
func Sqrt(x float64) (float64, error) {
	var z float64
	var a float64

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
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
