package main

import "fmt"

// we might want to close a channel once we are done using it

func main() {
	// in this example we will make a jobs channel to communicate between the work to be done and main
	// we will then close the channel when we are done
	jobs := make(chan int, 5)
	done := make(chan bool)

	// anon function for getting jobs
	go func() {
		for {
			// if more is false the channel is closed
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// sends three jobs to the jobs channel
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	// we then close the channel
	close(jobs)
	fmt.Println("sent all jobs")

	// await the worker using channel synchronization
	<-done

	// ok will be a bool if there are more jobs
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
