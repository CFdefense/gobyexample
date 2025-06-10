package main

import (
	"fmt"
	"slices"
)

func main() {

	// slices are a lightway dynamic view of an arrow
	// they will dynamically grow and are init with nil
	// lets make a new slice
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// if we want to make a new slice without init with nil
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// we can update some values with indexxing
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// we can also append which returns a new slice with the added value
	// this will dynamically auto resize the slice if needed
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// we can get a slice of a slice
	l := s[2:5]
	fmt.Println("sl1:", l)

	// slices up to and exlcuding index 5
	l = s[:5]
	fmt.Println("sl2:", l)

	// slices up from and including index 2
	l = s[2:]
	fmt.Println("sl3:", l)

	// create and initialize a sliuce
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// slices package contains some useful utility functions
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// can create dynamic 2d data structures with slices
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
