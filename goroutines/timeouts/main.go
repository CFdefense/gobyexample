package main

import (
	"fmt"
	"time"
)

// Timeouts are important for programs that connect to external resources or
//  that otherwise need to bound execution time.
// Implementing timeouts in Go is easy and elegant thanks to channels and select.

func main() {

	// For our example, suppose we’re executing an external call that
	// returns its result on a channel c1 after 2s.
	// Note that the channel is buffered, so the send in the goroutine is nonblocking.
	// This is a common pattern to prevent goroutine leaks in case the channel is never read.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// here we use select for the timeout check
	// if we are able to take the value from the buffered channel all is good
	// otherwise after some seconds we will timeout

	// since the above function takes 2 seconds we will timeout
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// in this example we extend the time to 2 seconds
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	// therefore if we now select we will get the value and avoid the timeout
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
