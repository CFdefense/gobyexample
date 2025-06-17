package main

import (
	"fmt"
	"time"
)

// Timers are for when you want to do something once in the future - tickers are for when you want to do something repeatedly at regular intervals.
func main() {

	// lets create a new ticker to tick at 500 ms and a channel for done
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// here well use a select to await the ticker values as they arrive
	// because this is a goroutine anon function it will run concurrently
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// tickers can also be stopped. Well stop it like this
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
