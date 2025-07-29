package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// we can test whether or not a string matches a regex pattern
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) // true

	// we can also compile the regex pattern and match based on the regex struct
	r, _ := regexp.Compile("p([a-z]+)ch")

	// we can perform that same match we did above
	fmt.Println(r.MatchString("peach")) // true

	// we can find a match in a string
	fmt.Println(r.FindString("peach punch")) // peach

	// this will find the first matches index range
	fmt.Println("idx:", r.FindStringIndex("peach punch")) // idx [0 5]

	// returns both full and submatches
	fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea]

	// will do the same as the above but return indexes
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

	// all variants for matching, -1 will return all matches
	fmt.Println(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]

	// finds all substring indexes
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1)) // all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	// find all with non-negative limits returned matches
	fmt.Println(r.FindAllString("peach punch pinch", 2)) // [peach punch]

	// can use Match on byte array
	fmt.Println(r.Match([]byte("peach"))) // true

	// MustCompile panics instead of returning an error, which makes it safer to use for global variables.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// can be used to replace substr matches
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>

	// can be used to edit some text
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out)) // a PEACH
}
