package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	A, B float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.A*v.A + v.B*v.B)
}

func (v *Vertex) Scale(f float64) {
	v.A = v.A * f
	v.B = v.B * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(5)
	fmt.Println("Abs of a triangle of sides 3 and 4 is", v.Abs())
}
