package main

import (
	"fmt"
	"math"
)

// /////////////////
// SHAPE
// /////////////////
type Shape interface {
	Area() float64
}

// /////////////////
// STRUCT
// /////////////////
type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

// /////////////////
// TYPE
// /////////////////
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// /////////////////
// FUNCTION
// /////////////////
func getArea(shape Shape) {
	fmt.Println(shape.Area())
}

func main() {
	r := Rectangle{Width: 7, Height: 8}
	c := Circle{Radius: 5}

	getArea(r)
	getArea(c)
}
