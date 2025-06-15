package main

import (
	"fmt"
	"time"
)

// We can use channels to synchronize execution across goroutines.
// Here’s an example of using a blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.

// This is the function we’ll run in a goroutine.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// The done channel will be used to notify another goroutine that this function’s work is done.
	done <- true
}

func main() {

	// Lets intialize the channel
	done := make(chan bool, 1)

	// Then we call the worker function
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	// If you removed the <- done line from this program,
	// the program would exit before the worker even started.
	<-done
}
