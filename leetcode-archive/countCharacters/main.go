package main

import (
	"fmt"
)

func canBeFormed(word string, chars string) bool {

	charsCount := make(map[rune]int)
	for _, r := range chars {
		charsCount[rune(r)]++
	}

	for _, r := range word {
		if val, found := charsCount[r]; found && val > 0 {
			charsCount[r]--
		} else {
			return false
		}
	}
	return true
}

func countCharacters(words []string, chars string) int {
	if len(words) < 1 {
		return 0
	}

	wLen := len(words)
	sum := 0

	for i:=0; i<wLen; i++ {
		word := words[i]
		if canBeFormed(word, chars) {
			sum += len(word)
		}
	}

    return sum
}

func main() {
	words := []string{"cat", "hat", "bt", "tree"} 
	chars := "atach"
	fmt.Printf("count chars %v %s -> %d\n", words, chars, countCharacters(words, chars))
}