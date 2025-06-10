package main

import (
	"fmt"
	"math"
)

// global constant example
const s string = "constant"

func main() {
	fmt.Println(s)

	// implied types and operations with constants
	const n = 500000
	const d = 3e20 / n

	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
