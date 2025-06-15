package main

import (
	"fmt"
	"time"
)

// 2A goroutine is a lightweight thread of execution

// lets define a normal function for example purposes
func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// normal synchronous function call
	f("direct")

	// if we use 'go' we can execute this function concurrently
	go f("goroutine")

	// we can also call anon functions this way
	// these two functions are now running concurrently.
	// for a better approach use waitgroups
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
