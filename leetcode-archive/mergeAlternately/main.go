package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	if (len(word1) < 1 || len(word1) > 100 || len(word2) < 1 || len(word2) > 100) {
		return ""
	}

	wLen1 := len(word1)
	wLen2 := len(word2)

	sword1 := []byte(word1)
	sword2 := []byte(word2)

	result := make([]byte, 0)

	for i,j := 0, 0; i<wLen1 || j<wLen2; i,j = i+1, j+1 {

		if i < wLen1 {
			result = append(result, sword1[i])
		}

		if j < wLen2 {
			result = append(result, sword2[j])
		}
	}

    return string(result)
}

func main() {
	w1 := "abcd1234"
	w2 := "pqrs"
	fmt.Printf("%s %s -> %s\n", w1, w2, mergeAlternately(w1,w2))
}