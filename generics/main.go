package main

import "fmt"

// generic types allow for multiple types to be passed into the same function
// lets take this example function for example
// this function requres a slice of type E, and a value E
// this ensures that we only execute logical operations such as '==' on comparable types
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// another example is say we have a list
// we want to implement some struct for this list but
// the catch is we want to have lists of many types, in order to avoid
// code replication we can specify the list can have any type
type List[T any] struct {
	head, tail *element[T]
}

// we must also do this with the lists embeded struct
type element[T any] struct {
	next *element[T]
	val  T
}

// here we define a method of the generic type
// we simlpy need to ensure we keep the type 'T' in place for all uses
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// heres another example of a method in which we use generics
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// notice how the compiler infers the types here and we dont need to explicitly say1
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	_ = SlicesIndex(s, "zoo")

	// we can create a list of lets say ints and add more ints
	// however we can also reuse the same List struct for a list of strings
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
