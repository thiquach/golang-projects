package main

import (
	"fmt"
)

func romanToInt(s string) int {

	if len(s) < 1 {
		return 0
	}

	romanMap := make(map[rune]int)
	romanMap['I'] = 1
	romanMap['V'] = 5
	romanMap['X'] = 10
	romanMap['L'] = 50
	romanMap['C'] = 100
	romanMap['D'] = 500
	romanMap['M'] = 1000

	sRunes := []rune(s)
	result := 0
	for i:=0; i<len(sRunes); i++ {
		rCurr, ok := romanMap[sRunes[i]]
		if !ok {
			fmt.Println("unreccognised roman character")
		} 

		if (i+1) < len(sRunes) {
			rNext, ok := romanMap[sRunes[i+1]]
			if !ok {
				fmt.Println("unreccognised roman character")
			} 
			if rNext > rCurr {
				result += (rNext - rCurr)
				i++
				continue
			}
		}

		result += rCurr
	}
	return result
}

func main() {
	// s := MCMXCIV 1994
	s := "XCIV" 	
	fmt.Printf("roman to int %s-> %d\n", s, romanToInt(s))
}