package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	A, B float64
}

func (v *Vertex) Scale(f float64) {
	v.A = v.A * f
	v.B = v.B * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.A*v.A + v.B*v.B)
}

func main() {
	p := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", p, p.Abs())
	p.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", p, p.Abs())
}
