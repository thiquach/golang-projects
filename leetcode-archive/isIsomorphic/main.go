package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sRunes := []rune(s)
	tRunes := []rune(t)
	sMap := make(map[rune]rune)
	tMap := make(map[rune]rune)

	for i:=0; i<len(sRunes); i++ {
		sr, ok := sMap[sRunes[i]]
		if ok {
			if sr != tRunes[i] {
				return false
			} 
		} else {
			sMap[sRunes[i]] = tRunes[i] 
		}

		tr, ok := tMap[tRunes[i]]
		if ok {
			if tr != sRunes[i] {
				return false
			} 
		} else {
			tMap[tRunes[i]] = sRunes[i] 
		}
	}
    return true
}

func main() {
	s := "add" 	// add & egg
	t := "egg"
	fmt.Printf("isomorphic %s %s -> %t\n", s, t, isIsomorphic(s, t))
}