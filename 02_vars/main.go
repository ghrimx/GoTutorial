package main

import "fmt"

// Globally declared variable (outside of main)
var globalvariable string = "I'm a global variable"

func main() {
	// Explore variable types

	// declaration and initialization
	var firstName string = "MyFirstName"

	// If an initializer is present, the type can be omitted
	var name = "MyName"

	// multiple declaration and initialization
	var str1, str2 string = "str1", "str2"

	// Short declaration, 'var' ignored
	string1, string2 := "first string", "second string"

	// Block declaration
	var (
		boolean bool   = true
		integer int    = 24
		MaxInt  uint64 = 1<<64 - 1
	)

	// Print out variable content
	fmt.Println(name, "\n", firstName, "\n", globalvariable, "\n", string1, "\n", string2, "\n", str1, "\n", str2)

	// Print out type and value
	// %T for type and %v for value
	fmt.Printf("Type: %T Value: %v \n", boolean, boolean)
	fmt.Printf("Type: %T Value: %v \n", integer, integer)
	fmt.Printf("Type: %T Value: %v \n", MaxInt, MaxInt)
}
