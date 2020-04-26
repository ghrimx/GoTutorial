package main

import "fmt"

// Main
func main() {
	fmt.Println("--- Functions ---")
	// To call: funcName(arguments)

	// No arguments no returns
	foo()

	// Call with argument
	bar("James")

	// Call with argument and return
	s1 := woo("Moneypenny")
	fmt.Println(s1)

	// Call with arguments and returns
	x, y := multiReturn("Ian", "Fleming")
	fmt.Println(x, y)

	// variadic parameter
	total := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("total", total)

	// Unfurling a slice
	sliceOfInt := []int{2, 3, 4, 5, 6, 7, 8, 9}
	anotherTotal := sum(sliceOfInt...)
	fmt.Println("anotherTotal", anotherTotal)

	fmt.Println("no argument to unlimited parameters function", sum())

	xoo("James")

	// Defer

} // end of main

// func (r receiver) identifier(parameters) (returns) { code }
// we define our func with parameters (if any)
// we call our func and pass in arguments (if any)
func foo() {
	fmt.Println("Hello from foo")
}

// function with parameter
func bar(s string) {
	fmt.Println("Hello", s)
}

// function returning something
func woo(s string) string {
	return fmt.Sprint("Hello from woo ", s)
}

// function returning multiple values
func multiReturn(firstName string, lastName string) (string, bool) {
	fmt.Println("function returning multiple values")
	a := fmt.Sprint(firstName, " ", lastName, `, says "Hello"`)
	b := false
	return a, b
}

// variadic parameter
// ... is a lexical elements - operator or delimiter
// the variadic paramter as to be the final one in the parameters list
func sum(x ...int) int {
	fmt.Println("variadic parameter")
	fmt.Printf("%v\t %T\n", x, x) // we passed individual value of type int and Go converted them into a slice
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println(i, v)
	}
	return sum
}

func xoo(s string, y ...int) {
	fmt.Printf("string: %v - Variadic parameter: %v \t %T \n", s, y, y)
}
