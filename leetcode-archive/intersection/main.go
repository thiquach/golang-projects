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

func intersection(nums1 []int, nums2 []int) []int {

	result := make([]int, 0, len(nums1) + len(nums2))
    for _, n := range nums2 {
        if slices.Contains(nums1, n) {
			result = append(result, n)
		}
    }
	
	return removeDuplicates(result)
}

func main() {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	result := intersection(nums1, nums2)

	for i := 0; i < len(result); i++ {
		fmt.Printf("%d\t", result[i])
		fmt.Println()
	}
}