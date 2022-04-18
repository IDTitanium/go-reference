package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangl struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangl) area() float64 {
	return math.Pi * r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	cicle := Circle{0, 0, 5}
	rect := Rectangl{10, 5}

	fmt.Printf("circle area: %f\n", getArea(cicle))
	fmt.Printf("rectangle area: %f\n", getArea(rect))
}
