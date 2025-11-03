package main

import (
	"fmt"
)

func countCharacters(words []string, chars string) int {
	if len(words) < 1 {
		return 0
	}

	charsCount := make([]int, 26)

	for _, r := range chars {
		charsCount[int(r - 'a')]++
	}

	currentCharsCount := make([]int, 26)
	sum := 0
	for _, word := range words {
		copy(currentCharsCount, charsCount)
		isValid := true

		for _, r := range word {
			currentCharsCount[r - 'a']--
			if (currentCharsCount[r - 'a'] < 0) {
				isValid = false
				break
			}
		}
		if isValid {
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