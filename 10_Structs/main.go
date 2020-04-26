package main

import (
	"fmt"
)

// main : Structs
// A struct is an composite data type. (composite data types, aka, aggregate data types, aka, complex data types). Structs allow us to compose together values of different types.

// Struc declaration
type person struct {
	first string
	last  string
	age   int
}

// Embedded struct
type secretAgent struct {
	person        // anonymous field & unqualified type name
	first  string // collision with first in person struct
	kill   bool
}

func main() {
	fmt.Println("--- Structs ---")
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	// Embedded struct
	fmt.Println("Embedded struct")
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32},
		first: "collision",
		kill:  true,
	}

	fmt.Println(p1, p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(sa1.person.first, sa1.first, sa1.last, sa1.age, sa1.kill)

	// Anonymous struct
	car := struct {
		brand, color string
		price        int
	}{
		brand: "Toyota",
		color: "red",
		price: 20000,
	}
	fmt.Println(car)
}
