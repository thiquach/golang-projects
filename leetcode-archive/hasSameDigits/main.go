package main

import (
	"fmt"
	"strings"
)

func hasSameDigits(s string) bool {
	if len(s) < 2 {
		return false
	}

	for len(s) > 2 {
		var builder strings.Builder
		builder.Grow(len(s) - 1) // allocate enough space

		for i := 1; i < len(s); i++ {
			sum := (s[i-1]-'0' + s[i]-'0') % 10
			builder.WriteByte(sum + '0')
		}

		s = builder.String()
	}

	return s[0] == s[1]
}

func main() {
	s := "50289"
	fmt.Println("has same digits? ", hasSameDigits(s))
}