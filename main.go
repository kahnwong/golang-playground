package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length, Height float64
}

type Triangle struct {
	Base, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return 2 * (r.Length + r.Height)
}

func (t Triangle) Area() float64 {
	return (0.5) * t.Base * t.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func getAttributes(name string, shape Shape) {
	fmt.Println(name)
	fmt.Printf("area: %f \n", shape.Area())
	fmt.Println("---------")
}

func main() {
	r := Rectangle{Length: 2, Height: 8}
	getAttributes("rectangle", r)

	t := Triangle{Base: 2, Height: 8}
	getAttributes("triangle", t)

	c := Circle{Radius: 4}
	getAttributes("circle", c)
}
