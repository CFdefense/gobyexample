package main

import (
	"fmt"
	"sync"
	"time"
)

// in order to implement functionality to 'wait' for all goroutines to finish:
// we must implement a 'waitgroup'

// here is our worker function
// it sleeps to mimic compute intensive work
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// here well initalize our waitgroup
	// the waitgroup will wait for all goroutines to finish
	// if the waitgroup is passed to a function pass as a reference
	var wg sync.WaitGroup

	// here well go and run multiple workers
	// we will add to the waitgroup each time
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		//
		go func() {
			// defer will execute after worker(i) returns
			// done will decrement the wg count
			defer wg.Done()
			worker(i)
		}()
	}

	// block until the wg is empty and all workers are done
	wg.Wait()

}
