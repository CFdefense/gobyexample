package main

import "fmt"

// go doesnt have direct enum support
// however enums can be created using existing language features
// lets start off by declaring our enum
type ServerState int

// the enum values are defined here as constants
// iota assigns successive integers for each constant
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// using a map we can map enum states to strings
// this can become cumbersome to implement
// should look into stringer and go:generate
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// an example of getting the string of the map
func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	// can pass the enum type into functions
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// here is an enum transition function, it takes in one enum value and returns another
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:

		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
