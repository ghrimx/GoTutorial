package main

import "fmt"

// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
// SolHO06lvl2 : Solution to ex 6

const (
	y1 = 2020 + iota // 2020 here iota equal 0
	y2 = y1 + iota   // 2021
	y3 = y1 + iota   // 2022
	y4 = y1 + iota   // 2023
)

func SolHO06lvl2() {
	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
	fmt.Println(y4)
}
