package main

import (
	"fmt"
	"time"
)

// We often want to execute Go code at some point in the future, or repeatedly at some interval.
// Go’s built-in timer and ticker features make both of these tasks easy.
func main() {

	// lets create a new timer - this one will wait 2 seconds
	timer1 := time.NewTimer(2 * time.Second)

	// we block the timer channel until it sends a value indicating the timer expired
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// lets create another timer
	timer2 := time.NewTimer(time.Second)

	// we then block the channel and await the timer to expire
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// however we can also stop the timer, this is done before the timer expires
	// so we see "Timer 2 Stopped"
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Can also use sleep like this
	time.Sleep(2 * time.Second)
}
