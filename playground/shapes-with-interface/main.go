package main

import (
	"fmt"
	"math"
)

type shapes interface {
	area() float64
	perim() float64
}

func measure(s shapes) {
	fmt.Printf("Shape %+v\n", s)
	fmt.Println("Area is", s.area())
	fmt.Println("Perimeter is", s.perim())
}

func main() {

	sq := Square{3}
	cir := Circle{5}
	measure(sq)
	measure(cir)
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return float64(s.side * s.side)
}

func (s Square) perim() float64 {
	return float64(s.side * 4)
}

func (c Circle) area() float64 {
	return float64(math.Pi * c.radius * c.radius)
}

func (c Circle) perim() float64 {
	return float64(2 * math.Pi * c.radius)
}
