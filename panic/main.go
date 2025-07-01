package main

// A panic typically means something went unexpectedly wrong.
// Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.
func main() {

	// this will terminate the program and display 'a problem'
	panic("a problem")

	// unreachable code due to panic above
	// _, err := os.Create("/tmp/file")
	// if err != nil {
	// 	panic(err)
	// }
}
