// Leetcode 66 - Plus One
package main

import (
	"fmt"
	"strconv"
)

func sliceToInt(s []int) int {
    res := 0
    op := 1
    for i := len(s) - 1; i >= 0; i-- {
        res += s[i] * op
        op *= 10
    }
    return res
}

func plusOne(digits []int) []int {
	if (len(digits) == 0) {
		return []int {}
	}

	value := sliceToInt(digits)
	value++

	s := strconv.Itoa(value)
	result := make([]int, 0, len(s)) 

	for _, char := range s {
		digit, _ := strconv.Atoi(string(char))
		result = append(result, digit)
	}

	return result
}

func main() {
	nums := []int{4,3,1,2}
	result := plusOne(nums)
	fmt.Println("result ", result)
}