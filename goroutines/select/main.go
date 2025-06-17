package main

import (
	"fmt"
	"time"
)

// Go’s select lets you wait on multiple channel operations.
// Combining goroutines and channels with select is a powerful feature of Go.

func main() {

	// lets select across two channels
	c1 := make(chan string)
	c2 := make(chan string)

	// these two channels will recieve a value after some time to simulate API calls
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// we’ll use select to await both of these values simultaneously, printing each one as it arrives.
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
