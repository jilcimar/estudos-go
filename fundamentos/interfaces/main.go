package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	height float64
	width  float64
}

type circle struct {
	radius float64
}

// Interface
type form interface {
	area() float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func calculateArea(f form) {
	fmt.Printf("A área da forma é %0.2f \n", f.area())
}

func main() {
	fmt.Println("Interfaces")

	r := rectangle{10, 15}
	calculateArea(r)

	c := circle{10}
	calculateArea(c)
}
