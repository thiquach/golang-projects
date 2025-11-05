package main

import (
	"fmt"
)

func rearrangeCharacters(s string, target string) int {

    if len(target) > len(s) {
        return 0
    }

	sCount := make([]int, 26)
	tCount := make([]int, 26)

	for _, r := range s {
		sCount[r - 'a']++
	}

	for _, r := range target {
		tCount[r - 'a']++
	}

	nCopy := 100
	for _, r := range target {
		x := float64(sCount[r - 'a'])
		y := float64(tCount[r - 'a'])
		a := int(x / y)
		if a < nCopy {
			nCopy = a
		}
	}
   return nCopy 
}

func main() {
	// s := "ilovecodingonleetcode" and target := code
	// s := "abbaccaddaeea" and aaaaa
	// wvu and tu
	// lgvhv and ly
	s := "lgvhv"
	target := "ly"
	fmt.Printf("reangeCharacters %s %s -> %d\n", s, target, rearrangeCharacters(s, target))
}