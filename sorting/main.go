package main

import (
	"fmt"
	"slices"
)

// Go’s slices package implements sorting for builtins and user-defined types. We’ll look at sorting for builtins first.
func main() {

	// initiaslize some array of string
	strs := []string{"c", "a", "b"}

	// Sorting functions are generic, and work for any ordered built-in type
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// now ints and we can see this also works
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// can also check for if its sorted
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
