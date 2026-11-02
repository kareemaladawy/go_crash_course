package main

import (
	"fmt"
	"math"
)

// An interface is a collection of method signatures
type Shape interface {
	getArea() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) getArea() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.getArea()
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{10, 5}

	fmt.Println(getArea(c))
	fmt.Println(getArea(r))
}
