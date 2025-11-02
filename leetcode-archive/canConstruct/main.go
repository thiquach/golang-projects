package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) < 1 || len(magazine) < 1 {
    	return false
	}

	magazineCount := make(map[rune]int)
	for _, r := range magazine {
		magazineCount[r]++
	}

	for _, r := range ransomNote {
		if val, ok := magazineCount[r]; ok {
			fmt.Printf("Value for %c exists: %d\n", r, val)
			if val > 0 {
				magazineCount[r]--
			} else {
				return false
			}
		} else {
			fmt.Println("Key does not exist.")
			return false
		}
	}

	return true
}

func main() {
	ransomNote := "aa" 
	magazine := "ab"
	fmt.Printf("canConstruct %s %s -> %t\n", ransomNote, magazine, canConstruct(ransomNote, magazine))
}