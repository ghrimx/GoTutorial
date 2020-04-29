package main

import "fmt"

// a value can be of more than one type
// if you have this method you are my type
type human interface {
	speak()
}

type myInterface interface {
	ThatMethod()
}

// interfaces: interfaces are used to do polymorphism
// empty interface (see ...interface{}) means unlimited number of value of type empty interface meaning "any value of any type"
func interfaces(h human) {
	fmt.Println("*** Interfaces")
	fmt.Println("I was passed into interfaces", h)
}

func ifYouHaveThatMethod(m myInterface) {
	fmt.Println("if you have that method", m)
}

func typeAssertion(h human) {
	switch h.(type) {
	case person:
		fmt.Println("you are a person", h.(person).firstName)
	case secretAgent:
		fmt.Println("you are a secre agent", h.(secretAgent).firstName)
	default:
		fmt.Println("unknown type")
	}
}
