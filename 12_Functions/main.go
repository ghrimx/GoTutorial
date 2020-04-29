package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	kill bool
}

// methods
// func (r receiver) identifier(parameters) (returns) {code}

func (p person) ThatMethod() {
	fmt.Println("you are my type")
}

func (s secretAgent) speak() {
	fmt.Println("I'm", s.firstName, s.lastName)
}

func (p person) speak() {
	fmt.Println("I'm", p.firstName, p.lastName)
}

// Main
func main() {

	p1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		kill: true,
	}
	fmt.Println(p1)

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
	defer zoo()
	postZoo()

	// Method
	fmt.Println("*** Methods")
	p1.speak()

	// Interfaces
	// see interfaces.go
	// person type has the method speak. Thanks to the interface person variables are also of type human
	p2 := person{
		firstName: "Homer",
		lastName:  "Simpson",
	}
	interfaces(p2) // p2 is of type person
	interfaces(p1) // p1 is of type secretAgent

	ifYouHaveThatMethod(p2)
	p2.ThatMethod()

	// type assertion
	// see interfaces.go
	typeAssertion(p1)
	typeAssertion(p2)

	fmt.Println(42)

	// Anonymous function
	// notice the parentheses after the func block
	fmt.Println("*** Anonymous function")
	func(x int) {
		fmt.Println(" Anonymous function ran", x)
	}(42)

	// func expression
	// assigning a function to a variable
	f := func() {
		fmt.Println("*** func expression")
		fmt.Println("my first func expression")
	}

	f()
	fmt.Printf("%T\n", f) // type func

	g := func(x int) {
		fmt.Println(x)
	}

	g(4242)

	// Returning a func from a func
	// Functions are types
	// when can this be usefull ? in web developement see http
	fmt.Println("*** Returning a func")
	fc := returnFunc() // this func return a func
	fmt.Printf("%T\n", fc)
	val := fc() // now it returns an int
	fmt.Println(val, fc())
	fmt.Println(returnFunc()())

	// Callback
	// Pass a func as an argument
	fmt.Println("*** Callback")
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	sumOfEvenNumber := even(sum, ii...)
	fmt.Println("sumOfEvenNumber:", sumOfEvenNumber)

} // end of main

//*************************************************
//*                FUNCTIONS                      *
//*************************************************

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
// it means infinite of
func sum(x ...int) int {
	fmt.Println("variadic parameter")
	fmt.Printf("%v\t %T\n", x, x) // we passed individual value of type int and Go converted them into a slice
	total := 0
	for i, v := range x {
		total += v
		fmt.Println(i, v)
	}
	return total
}

func xoo(s string, y ...int) {
	fmt.Printf("string: %v - Variadic parameter: %v \t %T \n", s, y, y)
}

func zoo() {
	fmt.Println("*** defer")
	fmt.Println("zoo")
}

func postZoo() {
	fmt.Println("postZoo")
}

// Returning a function
func returnFunc() func() int {
	return func() int { // anonymous func
		return 451
	}
}

// Callback
// even takes a function returning a int and assign it to f, even takes also an infinite number of int
// finally even returns an int
func even(f func(x ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...) // we unfurle yi
}
