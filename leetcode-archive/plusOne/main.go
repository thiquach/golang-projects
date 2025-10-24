// Leetcode 66 - Plus One
package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func main() {
	nums := []int{4,3,1,2}
	result := plusOne(nums)
	fmt.Println("result ", result)
}