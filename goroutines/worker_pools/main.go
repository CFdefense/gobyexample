package main

import (
	"fmt"
	"time"
)

// this is our 'worker' function it will run mutliple concurrent instances.
// this worker will recieve work from the 'work' chan and output in the 'results' chan
// we will sleep to emulate a compute intensive task
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// lets start by sending work to our workers
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// these three workers are started
	// they are initially blocked due to no jobs
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// then we send in 5 jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// close jobs as we have no more work
	close(jobs)

	// here we collect all of the results
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
