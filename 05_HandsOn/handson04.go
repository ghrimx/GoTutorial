package main

import (
	"fmt"
)

type (
	MyInt = int
)

var a MyInt

// SolHO04 : Solution to Hands-On exercice 4
func SolHO04() {
	fmt.Printf("%T %v \n", a, a)
	a = 42
	fmt.Println(x)
}
