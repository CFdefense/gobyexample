package main

import "fmt"

// we can use the for and range iterators to iterate through values recieved from a channel
func main() {

	// lets make a channel of type string and buffer size of 2
	queue := make(chan string, 2)

	// lets put some values in the channel and close it
	queue <- "one"
	queue <- "two"
	close(queue)

	// now we can use range to iterate over the values
	for elem := range queue {
		fmt.Println(elem)
	}
}
