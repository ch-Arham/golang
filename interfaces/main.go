package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectange struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectange) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r := rectange{width: 3, height: 4}
	c := circle{radius: 5}

	// shapes - slice of shape
	shapes := []shape{r, c}

	// loop through shapes
	for _, shape := range shapes {
		fmt.Println("Shape Area:", shape.area())
	}

	// fmt.Println("Rectangle Area:", r.area())
	// fmt.Println("Circle Area:", c.area())
}
