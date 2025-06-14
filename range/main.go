package main

import "fmt"

func main() {

	// range allos us to iterate over data structures

	// iterating over a slice
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// iterating through a range
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// iterating through a map's keys and values
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// can also just iterate over keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// iterate over unicode points
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
