package main

import "fmt"

// lets initialize the struct
type rect struct {
	width, height int
}

// this is the syntax for how we define a method in go
func (r *rect) area() int {
	return r.width * r.height
}

// methods can be defined for a pointer or by value
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// go automatically handles point to value conversions
	// we might want to use a reference to avoid copying the value
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
