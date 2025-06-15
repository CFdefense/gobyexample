package main

import (
	"errors"
	"fmt"
)

// in go we must handle potential errors
// lets define a function f which takes in an int
// this function then returns both int and error types
func f(arg int) (int, error) {
	if arg == 42 {

		// we can create an error like this
		return -1, errors.New("can't work with 42")
	}

	// we want to return nil in error type to indicate no error
	return arg + 3, nil
}

// A sentinel error is a predeclared variable that is used to
// signify a specific error condition.
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

// in this function we give an example of how we can wrap errors
// we wrap errors with higher-level errors for additional context
// in fmt we do this using %w
func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// here we call f and print a message based on the error response
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	// here we loop and use errors.Is to check for specific errors
	for i := range 5 {
		// lets call the function and check if theres an error
		if err := makeTea(i); err != nil {

			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}
