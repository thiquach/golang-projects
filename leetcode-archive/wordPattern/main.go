package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {

	pLen := len(pattern)
	sWordsArray := strings.Fields(s)
	sLen := len(sWordsArray)

	if sLen != pLen {
		return false
	}

	pMap := make(map[rune]string)
	sMap := make(map[string]rune)

	for i, p := range pattern {
		word := sWordsArray[i]
		if val, found := pMap[p]; found && val != word  {
			return false
		} 

		if val, found := sMap[word]; found && val != p {
			return false
		}
		
		pMap[p] = word
		sMap[word] = p
	}

    return true
}

func main() {
	s := "dog cat cat dog" 
	pattern := "abba"	
	fmt.Printf("words: %s, pattern: %s-> %t\n", s, pattern, wordPattern(pattern, s))
}