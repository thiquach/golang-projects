package main

import "fmt"

func typeCheck(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int - (%v %T)\n", i, i)
	case string:
		fmt.Printf("string (%v %T len%d)\n", i, i, len(v))
	default:
		fmt.Printf("unexpected type - (%v %T)\n", i, i)
	}
}

func main() {
	typeCheck(42)
	typeCheck("hello")
	typeCheck(true)
}
