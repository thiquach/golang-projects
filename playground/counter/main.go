// closures in Go
package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	nextCounter := counter()
	fmt.Println("nextCounter() ", nextCounter())
	fmt.Println("nextCounter()  ", nextCounter())

	fmt.Println()
	secondCounter := counter()
	fmt.Println("secondCounter()  ", secondCounter())
	fmt.Println("secondCounter()  ", secondCounter())
	fmt.Println("secondCounter()  ", secondCounter())
}