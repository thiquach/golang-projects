package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {

	sort.Ints(nums)
	for i:=len(nums) - 3; i>=0; i-- {
		if (nums[i] + nums[i+1] > nums[i+2]) {
			return nums[i]+nums[i+1]+nums[i+2]
		}
	}
	return 0
}

func main() {
	// []int{1, 2, 1, 10}
	// [3,2,3,4]
	nums := []int{3,2,3,4}
	fmt.Println("largestPerimeter -> ", largestPerimeter(nums))
}
