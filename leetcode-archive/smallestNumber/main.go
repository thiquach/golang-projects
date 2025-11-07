package main

import (
	"fmt"
	"math"
	"strconv"
)

func smallestNumber(n int) int {
	if n < 1 {
		return 0
	}

	binString := strconv.FormatInt(int64(n), 2)
	binStringLen := len(binString)
	fmt.Printf("bit len is %s %d ", binString, binStringLen)
	smallNunber := int(math.Pow(2, float64(binStringLen))) - 1
    return smallNunber
}

func main() {
	n := 10
	fmt.Printf("smallestNumber %d -> %d\n", n, smallestNumber(n))
}