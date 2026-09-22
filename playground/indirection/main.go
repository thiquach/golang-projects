package main

import (
	"fmt"
)

type Vertex struct {
	A, B float64
}

func (v *Vertex) Scale(f float64) {
	v.A = v.A * f
	v.B = v.B * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.A = v.A * f
	v.B = v.B * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)
	fmt.Println("v after scale 2 then scale 10", v)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println("p after scale 3 then scale 8", p)
}
