package main

import (
	"fmt"
)

func main() {
	// there are only three froms of for loop
	// have a look to Extended Backus-Naur Form (EBNF)

	sum := 0
	// for 'init statement'; 'condition'; 'post statement'
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("For loop -> Sum:", sum)

	// init and post statement are optional
	// in this situation, it is similar to a "while" loop
	n := 1
	for n < 10 {
		n += n
	}
	fmt.Println("While loop -> n:", n)

	// nested loop
	for i := 0; i <= 10; i++ {
		fmt.Printf("i: %d \n", i)
		for j := 0; j <= 3; j++ {
			fmt.Printf("\t j: %d \n", j)
		}
	}

	// infinite loop & break statement
	fmt.Println("Break statement:")
	x := 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done")

	// continue statement means "continue to the next iteration stop reading code and come back at the begining of the loop"
	// '%' gives the remainder (it is the modulus operator)
	fmt.Println("Continue statement:")
	y := 0
	// print only odd numbers till 100
	for {
		y++
		if y > 100 {
			break
		}
		if y%2 != 0 {
			continue
		}
		fmt.Println(y)

	}

	// print all numbers from 33 to 132 as ASCII unicode point and hex
	fmt.Println("mini challenge:")
	for z := 33; z <= 132; z++ {
		fmt.Printf("%d \t %#U \t %x \n", z, z, z)
	}

}
