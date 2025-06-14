package main

import (
	"fmt"
	"math"
)

// lets implement an interface for geometric shapes
type geometry interface {
	area() float64
	perim() float64
}

// lets also intialize a rect and circle
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to implement all the methods in the interface.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// we can also define methods to take in a type of interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// heres another
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	// lets intialize our structs
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// they can both use a method of the interface
	measure(r)
	measure(c)

	// we can pass them as interfaces to these methods
	detectCircle(r)
	detectCircle(c)
}
