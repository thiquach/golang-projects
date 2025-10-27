package main

import (
	"fmt"
	"strings"
)

func rotateString(s string, goal string) bool {

	if (len(s) < 1) || len(goal) < 1 || len(s) != len(goal) {
		return false
	}

	return strings.Contains(s + s, goal)
}

func main() {
	// "abcde" vs "cdeab" - true
	// "abcde" vs "abced" - false
	s := "abcde"
	goal := "cdeab"
	fmt.Printf("[ %s %s ] -> %t\n", s, goal, rotateString(s, goal))
}