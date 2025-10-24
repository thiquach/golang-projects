package main

import (
	"fmt"
	"slices"
)

func removeDuplicates[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := []T{}
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

func findDifference(nums1 []int, nums2 []int) [][]int {
    nums1Set := removeDuplicates(nums1)
	nums2Set := removeDuplicates(nums2)
	
    diff1 := make([]int, 0, len(nums1Set))
	diff2 := make([]int, 0, len(nums2Set))

    for _, n := range nums1Set {
        if !slices.Contains(nums2Set, n) {
			diff1 = append(diff1, n)
		}
    }

    for _, n := range nums2Set {
        if !slices.Contains(nums1Set, n) {
			diff2 = append(diff2, n)
		}
    }

	return [][]int{diff1, diff2}
}

func main() {
	nums1 := []int{1,2,3,3}
	nums2 := []int{1,1,2,2}
	result := findDifference(nums1, nums2)

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Printf("%d\t", result[i][j])
		}
		fmt.Println()
	}
}