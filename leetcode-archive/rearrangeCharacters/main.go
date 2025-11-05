package main

import (
	"fmt"
	"math"
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

	minFloor := 0
	min := 0
	for _, r := range target {
		x := float64(tCount[r - 'a'])
		y := float64(sCount[r - 'a'])
        if y == 0 {
            min = 0
            break
        }
		minFloor = int(math.Floor(y/x))
		if min == 0 {
			min = minFloor
		} else {
			min = int(math.Min(float64(min),float64(minFloor)))
		}
	}
   return min 
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