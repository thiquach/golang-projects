package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b - a
	}
}

func fib1() func() int {
	fibSequence := make(map[int]int)
	n := 0
	i := 0
	return func() int {
		i = n
		result := 0
		if i == 0 {
			result = 0
		} else if i == 1 {
			result = 1
		} else {
			result = fibSequence[i-2] + fibSequence[i-1]
		}
		fibSequence[i] = result
		n++
		return result
	}
}

func main() {
	fmt.Println("Fibonacci number using closure - Method 1")
	f := fibonacci()
	for range 10 {
		fmt.Println(f())
	}

	fmt.Println("Fibonacci number using closure - Method 2")
	fib := fib1()
	for range 10 {
		fmt.Println(fib())
	}
}
