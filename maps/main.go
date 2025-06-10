package main

import (
	"fmt"
	"maps"
)

func main() {

	// can create and intialize an empty map (hashmap) with make
	m := make(map[string]int)

	// can add key-value pairs
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// get a value by key
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// print the length of the map
	fmt.Println("len:", len(m))

	// delete some value using a key from map
	delete(m, "k2")
	fmt.Println("map:", m)

	// empty the map
	clear(m)
	fmt.Println("map:", m)

	// can get a bool of it is present in the map
	// maps always return a value
	// this can help with or_insert rust equiv
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// can intiialize a map with some key-value pairs
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// map also has these useful utility functions
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
