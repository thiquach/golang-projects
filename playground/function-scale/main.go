package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	A, B float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.A*v.A + v.B*v.B)
}

func Scale(v *Vertex, f float64) {
	v.A = v.A * f
	v.B = v.B * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println("Abs of hypotenuse of a right angle triangle of sides 3 and 4 after Scale 10 is", Abs(v))
}
