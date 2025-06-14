package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// in go the concept of a character is called a rune
	// a rune is stored in UTF-8
	// a rune is a integer representing a UTF-8

	// lets assign s the literal value of "hello" in thai
	// this literal is now UTF-8 encoded
	const s = "สวัสดี"

	// this will get the bytes of the string not the actual length
	fmt.Println("Len:", len(s))

	// we can loop through the bytes as such
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// this will give the correct 'char' length/amount but needs to decode the UTF-8
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// this is an example of how different runes take up multiple indices
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// or we can use the same loop but with decoding the rune
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

// you can pass a rnue to a function
func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
