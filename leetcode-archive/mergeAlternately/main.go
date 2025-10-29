package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	if (len(word1) < 1 || len(word1) > 100 || len(word2) < 1 || len(word2) > 100) {
		return ""
	}

	minLen := len(word1)
	if len(word2) < minLen {
		minLen = len(word2)
	}

	sword1 := []byte(word1)
	sword2 := []byte(word2)

	result := make([]byte, 0)

	for i:=0; i<minLen; i++ {
		result = append(result, sword1[i])
		result = append(result, sword2[i])
	}

	if len(word1) > minLen {
		for i:=minLen; i<len(word1); i++ {
			result = append(result, sword1[i])
		}
	} else if len(word2) > minLen {
		for i:=minLen; i<len(word2); i++ {
			result = append(result, sword2[i])
		}
	}

    return string(result)
}

func main() {
	w1 := "abcd1234"
	w2 := "pqrs"
	fmt.Printf("%s %s -> %s\n", w1, w2, mergeAlternately(w1,w2))
}