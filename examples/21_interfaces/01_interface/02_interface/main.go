package main

import "fmt"

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}

func (z square) getType() string  {
	return "This is a square"
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Println(z.area())
}

func main() {
	s := square{10}

	fmt.Printf("%T\n",s)

	info(s)
}
