package main

import "fmt"

// Go supports embedding of structs and interfaces to express a more seamless composition of types.
// lets staet of by defining a struct of base
type base struct {
	num int
}

// this method of base will return the bases num
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// we also have a container struct which is made up of a base struct
type container struct {
	base
	str string
}

func main() {

	// if we then initialize the container we must also initialize the embedding
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// this is interesting, we can directly access the embeddings values through the container
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// or we can do it through the embedding for readibility
	fmt.Println("also num:", co.base.num)

	// we can call methods of the embedding on the container
	fmt.Println("describe:", co.describe())

	// if we define an inteface here
	type describer interface {
		describe() string
	}

	// we can add that interface implementation onto another struct like this
	var d describer = co
	fmt.Println("describer:", d.describe())
}
