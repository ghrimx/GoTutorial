package main

import "fmt"

func main() {
	sum := 0
	// for 'init statement'; 'condition'; 'post statement'
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// init and post statement are optional
	// "while" loop
	n := 1
	for n < 10 {
		n += n
	}
	fmt.Println("n:", n)
}
