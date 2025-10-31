// Leetcode 2022 - Convert 1D Array Into 2D Array
package main

import (
	"fmt"
)

func maximumWealth(accounts [][]int) int {
    maxWealth := 0
	for i:=0; i<len(accounts); i++ {
		wealth := 0
		for j:=0; j<len(accounts[i]); j++ {
			wealth += accounts[i][j]
		}
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}
	return maxWealth
}

func main() {
	account := [][]int{{1, 2, 3}, {3,2,1}}
	fmt.Println("maximumWealth -> ", maximumWealth(account))
}