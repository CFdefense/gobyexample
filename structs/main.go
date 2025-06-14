package main

import "fmt"

// intialize a person struct with a name and age
type person struct {
	name string
	age  int
}

// here we create a constructor which returns a pointer to the object
func newPerson(name string) *person {
	// we can create the object like this
	p := person{name: name}

	// we can edit the object
	p.age = 42

	// derefence and return the address
	return &p
}

func main() {

	// we can pass in the values
	fmt.Println(person{"Bob", 20})

	// we can also specify the values
	fmt.Println(person{name: "Alice", age: 30})

	// ommited fields are zero-valued
	fmt.Println(person{name: "Fred"})

	// create a pointer to the new object
	fmt.Println(&person{name: "Ann", age: 40})

	// call the constructor
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	// we can have anon objects where we dont name the struct as its only used once
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
