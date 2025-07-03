package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

// there are tons of ways we can format strings in go
// below are some of the most common

func main() {

	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p) // print obj instance

	fmt.Printf("struct2: %+v\n", p) // will include struct field names

	fmt.Printf("struct3: %#v\n", p) // go syntax rep of the value passed

	fmt.Printf("type: %T\n", p) // print the type of a value

	fmt.Printf("bool: %t\n", true) // for printing booleans

	fmt.Printf("int: %d\n", 123) // base10 formatting

	fmt.Printf("bin: %b\n", 14) // binary representations

	fmt.Printf("char: %c\n", 33) // char corresponding to int

	fmt.Printf("hex: %x\n", 456) // hex encoding

	fmt.Printf("float1: %f\n", 78.9) // float decimal formatting

	fmt.Printf("float2: %e\n", 123400000.0) // scientific notation
	fmt.Printf("float3: %E\n", 123400000.0) // scientific notation

	fmt.Printf("str1: %s\n", "\"string\"") // string printing basic

	fmt.Printf("str2: %q\n", "\"string\"") // double quoted printed strings

	fmt.Printf("str3: %x\n", "hex this") // string as base16 hex

	fmt.Printf("pointer: %p\n", &p) // print representation of pointer

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345) // width of an integer, use a number after the % in the verb. By default the result will be right-justified and padded with spaces.

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45) // also width of floats

	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) // left justify with '-'

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b") // right justified strings

	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b") // left justify with '-'

	// Sprintf formats and returns a string without printing it anywhere.
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// You can format+print to io.Writers other than os.Stdout using Fprintf.
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
