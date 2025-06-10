package main

import "fmt"

// variadic functions allow us to pass an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// we can call such a function like this
	sum(1, 2)
	sum(1, 2, 3)

	// or like this
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
