// Leetcode 2022 - Convert 1D Array Into 2D Array
package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	wordsArray := strings.Fields(s)
    return len(wordsArray[len(wordsArray)-1])
}

func main() {
	words := "   fly me   to   the moon  "
	fmt.Println("lengthOfLastWord -> ", lengthOfLastWord(words))
}