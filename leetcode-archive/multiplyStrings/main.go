package main

import (
	"fmt"
	"math/big"
)

func multiply(nums1 string, nums2 string) string {

	n1 := new(big.Int)
	n2 := new(big.Int)

	n1, ok1 := n1.SetString(nums1, 10)
	n2, ok2 := n2.SetString(nums2, 10)

	if !ok1 || !ok2 {
		fmt.Println("Error converting string to big.Int")
		return ""
	}

	product := new(big.Int).Mul(n1, n2)
    return product.String()
}

func main() {
	nums1 := "123"  // 456 & 77
	nums2 := "456"
	fmt.Printf("multiply strings %s, %s -> %s\n", nums1, nums2, multiply(nums1, nums2))
}