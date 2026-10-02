package main

import (
	"strings"
	// run the following command for the package
	// go get golang.org/x/tour/wc
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var ss = strings.Fields(s)
	wMap := make(map[string]int)

	for _, v := range ss {
		wMap[v]++
	}

	return wMap
}

func main() {
	wc.Test(WordCount)
}
