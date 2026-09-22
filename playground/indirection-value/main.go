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

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.A*v.A + v.B*v.B)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
