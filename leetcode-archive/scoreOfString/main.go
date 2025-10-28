package main

import (
	"fmt"
	"math"
)

func scoreOfString(s string) int {
	if (len(s) < 2 || len(s) > 100) {
		return 0
	}

	sum := 0
	for i := 0; i < len(s)-1; i++ {
		d1, d2 := s[i], s[i+1]
		sum += int(math.Abs(float64(int(d1) - int(d2))))
	}
   return sum 
}

func main() {
	s := "hello"
	fmt.Printf("%s -> %d\n", s, scoreOfString(s))
}