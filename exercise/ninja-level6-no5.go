//https://go.dev/play/p/r-ySC46x-4l

// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	s1 := square{
		length: 5,
	}
	c1 := circle{
		radius: 12.5,
	}

	info(s1)
	info(c1)

}

