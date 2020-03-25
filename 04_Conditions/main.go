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
}
