package main

import (
	"fmt"
)

func shuffle(nums []int, n int) []int {

	result := make([]int, 0, 2*n)
	
	s1 := nums[:n]
	s2 := nums[n:]

	for i:=0; i<n; i++ {
		result = append(result, s1[i], s2[i])
	}
	
	return result
}

func main() {
	nums := []int{1,2,3,11,22,33}
	n := 3
	result := shuffle(nums, n)
	fmt.Println("result is ", result)
}