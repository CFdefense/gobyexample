package main

import "fmt"

// functions can return multiple return types
func vals() (int, int) {
	return 3, 7
}

func main() {

	// heres how we would call such a function
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
