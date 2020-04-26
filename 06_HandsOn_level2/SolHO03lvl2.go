package main

import "fmt"

const (
	a        = 42   // untyped constant
	b string = "ok" // typed constant
)

// SolHO03lvl2 : Solution ex 3
func SolHO03lvl2() {
	fmt.Printf("%v \t %T \n", a, a)
	fmt.Printf("%v \t %T \n", b, b)
}
