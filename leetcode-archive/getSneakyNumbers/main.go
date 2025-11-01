package main

import (
	"fmt"
)

func getSneakyNumbers(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	result := []int{}
	seen := make(map[int]bool)
	for _, n := range nums {
		if seen[n] {
			result = append(result, n)
		}
		seen[n] = true
	} 
	return result
}

func main() {
	nums := []int{0,1,1,0}
	fmt.Printf("getSneakynumber %v -> %v\n", nums, getSneakyNumbers(nums))
}