package main

import "fmt"

// interface{} is emplty interface, may hold values of any type
// interface{} can be replaced by any
func describe(i any) {
	fmt.Printf("(%v %T)\n", i, i)
}

func main() {

	var i any
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

}
