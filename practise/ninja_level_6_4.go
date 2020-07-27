package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type square struct {
	length float64
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (s square) area() float64 {
	return s.length * s.length
}
type shape interface {
	area() float64
}
func info (sh shape) {
	switch sh.(type) {
	case circle:
		fmt.Println("Circle area is: ", sh.area())
	case square:
		fmt.Println("Square area is: ", sh.area())
	}
}
func main() {
	c := circle{
		radius: 9.123,
	}
	s := square{
		length: 10.243,
	}
	info(c)
	info(s)
}
