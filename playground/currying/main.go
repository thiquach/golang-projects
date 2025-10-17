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
	curryAddThree := curryAdd(add, 3)
	fmt.Println("curryAddThree add 5 ->", curryAddThree(5))
	fmt.Println("curryAddThree add 11 ->", curryAddThree(11))
}