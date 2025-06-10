package main

import "fmt"

func main() {

	// type implied by compiler
	var a = "initial"
	fmt.Println(a)

	// can assign multiple at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// more examples of implied types
	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// create and assign operator :=
	f := "apple"
	fmt.Println(f)
}
