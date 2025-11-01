package main

import (
	"fmt"
	"strings"
)

func prefixCount(words []string, pref string) int {

	if len(pref) < 1 || len(words) < 1 {
		return 0
	}

	count := 0;
	for i:=0; i<len(words); i++ {
		if strings.HasPrefix(words[i], pref) {
			count++
		}
	}

    return count
}

func main() {
	words := []string{"attention", "attend", "go", "back"}
	prefix := "at"	
	fmt.Printf("prefix count %v-> %d\n", words, prefixCount(words, prefix))
}