package main

import "fmt"

func sumIntVariadic(nums ...int) int {
	t := 0
	for _, v := range nums {
		t += v
	}
	return t
}

func sumIntSlice(nums []int) int {
	t := 0
	for _, v := range nums {
		t += v
	}
	return t
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("The sum is:", sumIntVariadic(x...))
	fmt.Println("The sum is:", sumIntSlice(x))
}
