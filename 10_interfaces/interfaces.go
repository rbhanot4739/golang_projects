package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, breadth float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

func getArea(s Shape) float64 {
	return s.area()
}
func main() {
	//var s Shape
	c := Circle{10}
	r := Rectangle{10, 20}

	fmt.Println(getArea(c))
	fmt.Println(getArea(r))
}
