package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// another way we can sync go program state across channel is with atomic counters

func main() {

	// atomic integer type to represent the timer
	var ops atomic.Uint64

	// a waitgroup to help wait for all goroutines to finish
	var wg sync.WaitGroup

	// add 50 goroutines that each increment the timer 1000 times
	for range 50 {
		wg.Add(1)

		go func() {
			for range 1000 {

				ops.Add(1)
			}

			wg.Done()
		}()
	}

	// wait till then all finish
	wg.Wait()

	// lets see the total value: 50000
	// had we not done this we would see differing values
	fmt.Println("ops:", ops.Load())
}
