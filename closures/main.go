package main

import "fmt"

// anonymous functions act as closures
// the returned function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// this variable holds an anonymous function
	nextInt := intSeq()

	// we can then call the function here
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
