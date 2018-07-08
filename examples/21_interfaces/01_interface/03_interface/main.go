package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type otherShape interface {
	area() float64
	getName() string
}

type square struct {
	side float64
}

// another shape
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

// which implements the shape interface
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Println(z.area())
}

func extendedInfo(os otherShape)  {
	fmt.Println(os.area())
	fmt.Println(os.getName())
}

func main() {
	s := square{10}

	c := circle{5}
	info(s)
	info(c)
}
