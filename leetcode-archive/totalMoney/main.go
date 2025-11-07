// Leetcode 2022 - Convert 1D Array Into 2D Array
package main

import (
	"fmt"
)

func totalMoney(n int) int {
	if n < 0 {
		return 0
	}

	weekStart := 1
	totalMoney := 0

	for i := 1; i <= n; i++ {
		totalMoney += weekStart + (i-1)%7

		if i%7 == 0 {
			weekStart++
		}
	}
	return totalMoney
}

func main() {
	n := 20
	fmt.Println("totalMoney -> ", totalMoney(n))
}