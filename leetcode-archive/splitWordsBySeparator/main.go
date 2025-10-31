package main

import (
	"fmt"
	"strings"
)

func splitWordsBySeparator(words []string, separator byte) []string {

	results := []string{}

    for _, w := range words {
		for _, p := range strings.Split(w, string(separator)) {
			if p != "" { 
				results = append(results, p)
			}
		}
	}

    return results
}

func main() {
	words := []string{"one.two.three", "four.five.six"}
	separator := byte('.')
	results := splitWordsBySeparator(words,separator)
	fmt.Printf("split strings by separator %v \n", results)
}