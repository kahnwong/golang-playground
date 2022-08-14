package main

import (
	"fmt"
)

// /////////////////
// SHAPE
// /////////////////
type Cost interface {
	Cost() float64
}

// /////////////////
// TYPE: LAMBDA
// /////////////////
type Lambda struct {
	Requests int
	Duration float64
}

func (l Lambda) Cost() float64 {
	return 0.0000002 * float64(l.Requests) * l.Duration
}

// /////////////////
// TYPE: EC2
// /////////////////
type EC2 struct {
	InstancePrice float64
	Duration      float64
}

func (e EC2) Cost() float64 {
	return 0.0000002 * float64(e.InstancePrice) * e.Duration
}

// /////////////////
// TYPE: ELB
// /////////////////
type ELB struct {
	Duration float64
}

func (e ELB) Cost() float64 {
	return 0.0000002 * 0.008 * e.Duration
}

// /////////////////
// FUNCTION
// /////////////////
func getCost(cost Cost) {
	fmt.Println(cost.Cost())
}

func main() {
	l := Lambda{Requests: 1000, Duration: 100}
	getCost(l)

	e := EC2{InstancePrice: 0.01, Duration: 100}
	getCost(e)

	elb := ELB{Duration: 100}
	getCost(elb)
}
