package main

import (
	"fmt"
	"strings"
)

func sumAndMod10(s string) string {
	if (len(s) < 2) {
		return s
	}

	var builder strings.Builder
	builder.Grow(len(s) - 1)

	for i:=0; i<len(s) - 1; i++ {
		d1 := s[i] - '0'
		d2 := s[i+1] - '0'
		val := (int(d1) + int(d2)) % 10
		builder.WriteByte(byte(val) + '0')
	}

	return builder.String()
}

func hasSameDigits(s string) bool {
	
	resultString := sumAndMod10(s)
	for len(resultString) > 2 {
		resultString = sumAndMod10(resultString)
	}

	return (resultString[0] == resultString[1])
}

func main() {
	s := "50289"
	fmt.Println("has same digits? ", hasSameDigits(s))
}