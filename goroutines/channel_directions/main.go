package main

import "fmt"

// When using channels as function parameters, you can specify if a channel
// is meant to only send or receive values. This specificity increases the type-safety of the program.

// recieve syntax:
// name chan<- type

// send syntax:
// name <-chan type

// function to add a message to a channel
// this channel parameter can only recieve
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// this other function accepts one parameter channel for recieving
// and it accepts another one for sending
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// lets create our channels
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// we can use ping to add to the channel
	ping(pings, "passed message")

	// then we pong with both
	pong(pings, pongs)

	// then we print a recieved value from the channel
	fmt.Println(<-pongs)
}
