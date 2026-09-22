package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	A, B float64
}

func (v Vertex) Hypotenuse() float64 {
	return math.Sqrt(v.A*v.A + v.B*v.B)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("Hypotenuse of a right angle triangle of sides 3 and 4 is", v.Hypotenuse())
}
