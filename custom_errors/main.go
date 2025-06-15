package main

import (
	"errors"
	"fmt"
)

// it is possible to use custom types as erors using Error()
// lets begin by declaring our custom type
// we usually want to end a custom error type with 'Error'
type argError struct {
	arg     int
	message string
}

// here well implement the error interface on our type
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// lets define a function which will return the error if it finds it
func f(arg int) (int, error) {
	if arg == 42 {

		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// call the function and get back value, error
	_, err := f(42)

	// now we can check for a match for a specific error type
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
