package main

import (
	"fmt"
	"math/big"
)

func addStrings(nums1 string, nums2 string) string {
	n1 := new(big.Int)
	n2 := new(big.Int)
	n1, ok1 := n1.SetString(nums1, 10)
	n2, ok2 := n2.SetString(nums2, 10)
    
	if (!ok1 || !ok2) {
		fmt.Println("Error converting string to big.Int")
		return ""
	}
	
	sum := new(big.Int).Add(n1, n2)
    return sum.String()
}

func main() {
	nums1 := "3876620623801494171"  // 456 & 77
	nums2 := "6529364523802684779"
	fmt.Printf("add strings %s, %s -> %s\n", nums1, nums2, addStrings(nums1, nums2))
}