package main

import "fmt"

// defer calling the function until the end of the program in the order of LIFO - Last In First Out
func main() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	fmt.Println(0)
}
