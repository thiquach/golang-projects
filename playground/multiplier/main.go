// Function returning a function
package main

import "fmt"

func multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func main() {

	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println("Double of 5 is ", double(5))
	fmt.Println("Triple of 3 is ", triple(3))
}