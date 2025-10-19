// Full currying in Go
package main

import "fmt"

func curryMultiply(x int) func(y int) func(z int) int {
	return func(y int) func(z int) int {
		return func(z int) int {
			return x * y * z
		}
	}
}

func main() {
	multiplyBy2 := curryMultiply(2)
	multiplyBy2And3 := multiplyBy2(3)
	result := multiplyBy2And3(5)

	fmt.Printf("Result of curryMultiply(2)(3)(5): %d\n", result)

	fmt.Printf("multiplyBy2And3(10): %d\n", multiplyBy2And3(10))
	fmt.Printf("multiplyBy2(4)(6): %d\n", multiplyBy2(4)(6))
}