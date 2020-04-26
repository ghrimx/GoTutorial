package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

// SolHO01lvl5 : Solution ex 1
func SolHO01lvl5() {
	fmt.Println("SolHO01lvl5:")
	p1 := person{
		firstName: "Mickael",
		lastName:  "Phelps",
		favIceCream: []string{
			"vanilla",
			"chocolate",
			"coconut",
		},
	}

	p2 := person{
		firstName: "Homer",
		lastName:  "Simpson",
		favIceCream: []string{
			"chocolate",
			"rum and coke",
			"Duff",
		},
	}

	fmt.Println(p1.firstName, p1.lastName)
	for i, v := range p1.favIceCream {
		fmt.Println(i, v)
	}

	fmt.Println(p2.firstName, p2.lastName)
	for i, v := range p2.favIceCream {
		fmt.Println(i, v)
	}
}

// SolHO02lvl5 : Solution ex 2
func SolHO02lvl5() {
	fmt.Println("SolHO02lvl5:")

	p1 := person{
		firstName: "Mickael",
		lastName:  "Phelps",
		favIceCream: []string{
			"vanilla",
			"chocolate",
			"coconut",
		},
	}

	p2 := person{
		firstName: "Homer",
		lastName:  "Simpson",
		favIceCream: []string{
			"chocolate",
			"rum and coke",
			"Duff",
		},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k, v := range m {
		fmt.Println(k, v.firstName)
		for j, val := range v.favIceCream {
			fmt.Println(j, val)
		}
	}

}

// SolHO03lvl5 : Solution ex 3
func SolHO03lvl5() {
	fmt.Println("SolHO03lvl5")
	c1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		fourWheel: true,
	}

	c2 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		luxury: true,
	}
	fmt.Println(c1, c2)
	fmt.Println(c1.doors, c1.fourWheel, c2.doors, c2.luxury)
}

// SolHO04lvl5 : Solution ex 4
func SolHO04lvl5() {
	fmt.Println("SolHO04lvl5")

	laptop := struct {
		brand         string
		cores, memory int
	}{
		brand:  "Toshiba",
		cores:  4,
		memory: 16,
	}

	fmt.Println(laptop)

	// OR

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}

// main : HandsOn level 5
func main() {
	fmt.Println("--- HandsOn level 5")

	SolHO01lvl5()

	SolHO02lvl5()

	SolHO03lvl5()

	SolHO04lvl5()

	// Try
	// map[int][]struct{} -> https://stackoverflow.com/questions/40254007/go-equivalent-of-pythons-dictionary
}
