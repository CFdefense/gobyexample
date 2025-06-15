package main

import (
	"fmt"
	"iter"
	"slices"
)

// lets implement the same generic List from 'generics'
type List[T any] struct {
	head, tail *element[T]
}

// and lets also implement its embedded struct
type element[T any] struct {
	next *element[T]
	val  T
}

// simple push methodfor the list
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// here instead of the 'generics' implementation of ALL
// we will use an all method implementation which returns an iterator
// The iterator function takes another function as a parameter, called yield by convention
// (but the name can be arbitrary). It will call yield for every element we want to iterate
// over, and note yield’s return value for a potential early termination.
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {

		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// heres another example using iterators for fibonnaci
// iterators dont require an underlying data structure
// this will continue until yield results in a false
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// now since list all returns an iterator we
	// can use it in a typical range loop
	for e := range lst.All() {
		fmt.Println(e)
	}

	// the method collect for slices takes an iterator
	// it then collects all of its values into a slice
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	// this will continue until false inside the fib
	for n := range genFib() {

		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
