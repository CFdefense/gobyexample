package main

import "fmt"

// pass by value
func zeroval(ival int) {
	ival = 0
}

// pass by refernece
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	// intially set i to 1
	i := 1
	fmt.Println("initial:", i)

	// we pass it by value to the function
	zeroval(i)
	fmt.Println("zeroval:", i)

	// then pass it by reference
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// we can print its memory address by dereferecing
	fmt.Println("pointer:", &i)
}
