package main

import (
	"fmt"
)

func numberOfSteps(num int) int {
    steps := 0
    for num > 0 {
        if num % 2 == 0 {
            num /= 2
        } else {
            num--
        }
        steps++
    }
    return steps
}

func main() {
	n := 14
	fmt.Println("Number of steps ", numberOfSteps(n))
}