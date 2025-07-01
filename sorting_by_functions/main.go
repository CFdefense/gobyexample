package main

import (
	"cmp"
	"fmt"
	"slices"
)

// Sometimes we’ll want to sort a collection by something other than its natural order.
// For example, suppose we wanted to sort strings by their length instead of alphabetically. Here’s an example of custom sorts in Go.
func main() {
	// lets say we have some array of strings
	fruits := []string{"peach", "banana", "kiwi"}

	// We implement a comparison function for string lengths. cmp.Compare is helpful for this.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// well then use 'SortFunc' to sort by this function
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// lets intialize a struct for a person
	type Person struct {
		name string
		age  int
	}

	// and then make some people objects
	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}

	// using 'SortFunc' again we can implement an anon function to sort this struct by
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
