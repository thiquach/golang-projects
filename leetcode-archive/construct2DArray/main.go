// Leetcode 2022 - Convert 1D Array Into 2D Array
package main

import (
	"fmt"
)

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < len(original); i++ {
		row := i / n
		col := i % n
		result[row][col] = original[i]
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	result := construct2DArray(nums, 2, 3)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[0]); j++ {
			fmt.Printf("%d\t", result[i][j])
		}
		fmt.Println()
	}
}