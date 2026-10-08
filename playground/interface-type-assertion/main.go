package main

import "fmt"

// Type assertion to check the type in the empty interface
func main() {
	var i any = "hello"

	s := i.(string)
	fmt.Println(s)

	f, ok := i.(float64)
	fmt.Println(f, ok)
	if ok == false {
		panic("unexpected type")
	}

}
