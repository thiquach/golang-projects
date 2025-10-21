// arrays and slices in Go
package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf(" %s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	primes := [6]int{2,3,5,7,11,13}
	fmt.Println("primes ", primes)

	var p = primes[1:4]
	fmt.Println("primes[1:4]", p)

	p = primes[0:]
	fmt.Println("primes[0:] ", p)

	p = primes[:2]
	fmt.Println("primes[:2] ", p)

	p = primes[1:]
	fmt.Println("primes[1:] ", p)

	p = primes[:]
	fmt.Println("primes[:] ", p)

	a := make([]int, 5)
	printSlice("Create slice using make([]int, 5) ", a)

	b := make([]int, 0, 5)
	printSlice("Create slice using make([]int, 0, 5) ", b)

	b = b[:cap(b)]
	printSlice(" b[:cap(b)],  ", b)

	b = b[1:]
	printSlice(" b[1:], ", b)
}