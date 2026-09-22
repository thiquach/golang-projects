package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return float64(s.side * s.side)
}

func (s Square) circum() float64 {
	return float64(s.side * 4)
}

func (c Circle) area() float64 {
	return float64(math.Pi * c.radius * c.radius)
}

func (c Circle) circum() float64 {
	return float64(2 * math.Pi * c.radius)
}

func main() {

	sq := Square{3}
	fmt.Printf("Square: %+v\n", sq)
	fmt.Printf("Square: area %f\n", sq.area())
	fmt.Printf("Square: circum %f\n", sq.circum())

	cir := Circle{5}
	fmt.Printf("Circle: %+v\n", cir)
	fmt.Printf("Circle: area %f\n", cir.area())
	fmt.Printf("Circle: circum %f\n", cir.circum())
}
