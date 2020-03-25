package main

import (
	"fmt"
	"runtime"
)

// Globally declared variable (outside of main)
var globalvariable string = "I'm a global variable"
var x bool

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
	fmt.Println(x) // zero-value of a bool is false

	// runtime constant
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	// String
	s := "Im a string"
	b := "I'm another string"
	t := `I'm a sentence
		  on two line`
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(t)

	// conversion of a string character to byte
	// behind the scene a string is a slice of byte
	fmt.Println("Slice of bytes")
	bs := []byte(s)
	fmt.Printf("%v \t %T \n", bs, bs)

	// loop over the string and print out the hex value of each letter
	fmt.Println("Print the index and the unicode value of each character")
	for i := 0; i < len(s); i++ {
		fmt.Printf("index %d : %#U \n", i, s[i])
	}

	fmt.Println("Print the index and the byte value of each character")

	for i, v := range s {
		fmt.Printf("index %d : %v \n", i, v)
	}

	fmt.Println("Print out index and the hex value")

	for i, r := range s {
		fmt.Printf("index %d : %#x \n", i, r)
	}

}
