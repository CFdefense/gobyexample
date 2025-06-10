package main

import "fmt"

func main() {

	// single condition for statement
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// example of intial/condition/after for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// can iterate with a range
	for i := range 3 {
		fmt.Println(i)
	}

	// no whiles in go we use for with no conditions
	for {
		fmt.Println("loop")
		break
	}

	// continue to go to next loop iteration
	for i := range 6 {
		if i%2 == 0 {
			continue
		}
	}
}
