package main

import "fmt"

// Globally declared variable
var globalvariable string = "I'm a global variable"

func main() {
	// Explore variable types

	// string
	var firstName string = "MyFirstName" // string keywords is not required as it is implicitly declared as string at assignement stage
	var name = "MyName"

	// in-line multiple assignment
	string1, string2 := "first string", "second string"

	// Print out variable content
	fmt.Println(name, "\n", firstName, "\n", globalvariable, "\n", string1, "\n", string2)

}
