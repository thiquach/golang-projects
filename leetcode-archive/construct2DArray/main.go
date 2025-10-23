// Leetcode 2022 - Convert 1D Array Into 2D Array
package main

import (
	"fmt"
)

func construct2DArray(original []int, r int, c int) [][]int {
	if len(original) != r*c {
		return [][]int{}
	}

	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}

	for i := 0; i < len(original); i++ {
		row := i / c
		col := i % c
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