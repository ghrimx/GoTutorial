package main

import (
	"fmt"
	"math"
)

// Simple if condition
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x)) // convert into string
}

// if condition and short statement
// if statement; condition
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// if else condition
func custpow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, lim) // 'v' is available in the else block
	}
	return lim
}

func custpowb(x, n, lim float64) float64 {
	v := math.Pow(x, n)
	if v > lim {
		fmt.Printf("%g >= %g \n", v, lim) // 'v' is available in the else block
		v = lim
	}
	return v
}

func custpowc(x, n, lim float64) float64 {
	v := math.Pow(x, n)
	if v < lim {
		return v
	}
	fmt.Printf("%g >= %g \n", v, lim)
	return lim
}

func main() {
	fmt.Println(
		sqrt(2),
		sqrt(-4),
		"pow:", pow(3, 2, 10),
		"custpow: ", custpow(3, 3, 20),
		"custpowb", custpowb(3, 3, 20),
		"custpowb_1", custpowb(3, 3, 40),
		"custpowc", custpowc(3, 3, 40),
		"custpowc", custpowc(3, 3, 20),
	)

	// ----- from Todd McLeod udemy course ----- //
	fmt.Println("From Tood McLeod")
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
	if 2 == 2 {
		fmt.Println("005")
	}
	if !(2 == 2) {
		fmt.Println("006")
	}
	if 2 != 2 {
		fmt.Println("007")
	}
	if !(2 != 2) {
		fmt.Println("008")
	}

	// short declaration in if statement
	if w := 42; w == 42 {
		fmt.Println(w == 42)
	}

	// else if
	x := 42
	if x == 40 {
		fmt.Println("x is ", x)
	} else if x == 41 {
		fmt.Println("x is ", x)
	} else {
		fmt.Println("x is", x)
	}

	// switch
	switch {
	case false:
		fmt.Println("case : not print")
	case (2 == 4): // false
		fmt.Println("case 001")
	case (3 == 3): // true
		fmt.Println("case 002")
		fallthrough // continue - not recommended to use this
	case (6 == 6): // true
		fmt.Println("case 003")
	default:
		fmt.Println("Print this if all cases are false")
	}

	// switch with multiple condition
	switch "c" {
	case "a", "c", "d":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	default:
		fmt.Println("default")
	}

	// default can be empty
	switch "c" {
	case "a", "c", "d":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	default:
	}

	// Logical operators
	// && -> AND
	// || -> OR
	// !  -> NOT
	fmt.Println("Logical operators:")
	p := true
	q := false
	if p && q {
		fmt.Println("&&")
	}
	if p || q {
		fmt.Println("||")
	}
	if p && !q {
		fmt.Println("!", !q)
	}

	fmt.Printf("true && true : %v \n", true && true)
	fmt.Printf("true && false : %v \n", true && false)
	fmt.Printf("true || true : %v \n", true || true)
	fmt.Printf("true || false : %v \n", true || false)
	fmt.Printf("!true: %v \n", !true)
	fmt.Printf("!false: %v \n", !false)

} // End of main
