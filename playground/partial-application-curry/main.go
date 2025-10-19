// currying in Go
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func curryAdd(f func(int, int) int, x int) func(y int) int {
	return func(y int) int {
		return f(x, y)
	}
}

func main() {
	add3 := curryAdd(add,3)
	result := curryAdd(add, 3)(5)
	
	fmt.Println("Result of curryAdd(add,3)(5) ->", result)
	fmt.Println("result of add3(12) ->", add3(12))
}