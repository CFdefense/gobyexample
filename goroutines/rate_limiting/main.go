package main

import (
	"fmt"
	"time"
)

func main() {

	// lets make a channel for our incoming requests
	// welcome have a buffer of 5 and add 5 requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// next we initialize a ticker at 200ms
	// this will be the rate limit regulator
	limiter := time.Tick(200 * time.Millisecond)

	// by blocking on the limiter we only allow 1 req per 200ms
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// we can also allow short bursts of requests
	// we accomplish this by buffering our limiter channel
	burstyLimiter := make(chan time.Time, 3)

	// begin by sending 3 requests
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// then every 200 ms we try to add a request up to max buffer of 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// we then add 5 more requests
	// the first 3 requests will be sent as a burst
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
