package main

import (
	"fmt"
	"strings"
)

func isVowel(r byte) bool {
	lowerR := strings.ToLower(string(r))

	switch lowerR {
	case "a", "e", "i", "o", "u":
		return true
	default:
		return false
	}
}

func reverseVowels(s string) string {
	if len(s) < 1 {
		return ""
	}

	sb := []byte(s) 
	left := 0
	right := len(s) - 1

	for left < right {
		for left < right && !isVowel(sb[left]) {
			left++
		}

		for left < right && !isVowel(sb[right]) {
			right--
		}

		if (sb[left] != sb[right]) {
			temp := sb[left]
			sb[left] = sb[right]
			sb[right] = temp

		}
		left++
		right--
	}
	return string(sb)
}

func main() {
	s := "IceCreAm"
	// s := "leetcode"
	fmt.Printf("%s -> %s\n", s, reverseVowels(s))
}