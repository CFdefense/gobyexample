package main

import "fmt"

// typical function definition with parameters and return type
func plus(a int, b int) int {

	return a + b
}

// can omit the param type if same
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// call function as a declare and initialize
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	// or just callon an existing var
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
