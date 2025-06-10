package main

import "fmt"

func main() {

	// initialize an array of 5 elements of type int
	var a [5]int
	fmt.Println("emp:", a)

	// set index 4 to 100
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	// can initialize with values
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// dont even need to pass the size
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// if you specify the index with :, the elements in between will be zeroed. ([100 0 0 400 500])
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// lets iterate through a 2d array
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// can also initialize 2d arrays with values
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
