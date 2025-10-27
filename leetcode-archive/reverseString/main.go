package main

import "fmt"

func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}

	left := 0
	right := len(s) - 1

	for left < right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
}

func main() {
	byteSlice := []byte{72, 101, 108, 108, 111}
	fmt.Printf("reverse %s ", string(byteSlice))
	reverseString(byteSlice)
	// for the reversed string - print s inside reverse() function
}