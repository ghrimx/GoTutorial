package main

import (
	"fmt"
)

// SolHO01lvl4 : Hands-on ex 1
func SolHO01lvl4() {
	a := [5]int{42, 43, 44, 45, 46}

	for i, v := range a {
		fmt.Printf("value: %v\t index: %v\n", v, i)
	}
	fmt.Printf("Type: %T\n", a) // size is part of the type
}

// SolHO02lvl4 : Hands-on ex 2
func SolHO02lvl4() {
	s := []int{47, 48, 49, 50, 51, 52, 53, 54, 55, 56}
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("Type: %T\n", s)
}

// SolHO03lvl4 : Hands-on ex 3
func SolHO03lvl4() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	s1 := s[:5]
	s2 := s[5:]
	s3 := s[2:7]
	s4 := s[1:6]
	fmt.Println("s1", s1)
	fmt.Println("s2", s2)
	fmt.Println("s3", s3)
	fmt.Println("s4", s4)
}

// SolHO04lvl4 : Hands-on ex 4
func SolHO04lvl4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	// for _, v := range y {
	// 	x = append(x, v)
	// }
	x = append(x, y...) // append the slice x with slice y in one statement
	fmt.Println(x)
}

// SolHO05lvl4 : Hands-on ex 5
func SolHO05lvl4() {
	z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	w := append(z[:3], z[6:]...)
	fmt.Println(w)
}

// SolHO06lvl4 : Hands-on ex 6
func SolHO06lvl4() {
	us := make([]string, 50, 50)
	// us = []string{"string1","string2",...}
	us = append(us, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println("len", len(us), "cap", cap(us))

	for i := 0; i < len(us); i++ {
		fmt.Println(i, us[i])
	}
}

// SolHO07lvl4 : Hands-on ex 7
func SolHO07lvl4() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	t := [][]string{jb, mp}
	for _, v := range t {
		for _, k := range v {
			fmt.Println(k)
		}
	}
}

// SolHO08lvl4 : Hands-on ex 8
func SolHO08lvl4() {
	// compostie literale is not mandatory when declaring a map
	d := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	// Hands-on ex 9
	d["My_Name"] = []string{"pastas", "Golang", "Data science"}
	// Hands-on ex 9 - end

	// Hands-on ex 10
	delete(d, "no_dr")
	// Hands-on ex 10 - end

	for _, v := range d {
		fmt.Printf("%T\n", v)
		for i, e := range v {
			fmt.Printf("%T\n", e)
			fmt.Println(i, e)
		}
	}
}

// MAIN
func main() {
	// Hands-on ex 1
	fmt.Println("Hands-on ex 1")
	SolHO01lvl4()

	// Hands-on ex 2
	fmt.Println("Hands-on ex 2")
	SolHO02lvl4()

	// Hands-on ex 3
	fmt.Println("Hands-on ex 3")
	SolHO03lvl4()

	// Hands-on ex 4
	fmt.Println("Hands-on ex 4")
	SolHO04lvl4()

	// Hands-on ex 5
	fmt.Println("Hands-on ex 5")
	SolHO05lvl4()

	// Hands-on ex 6
	fmt.Println("Hands-on ex 6")
	SolHO06lvl4()

	// Hands-on ex 7
	fmt.Println("Hands-on ex 7")
	SolHO07lvl4()

	// Hands-on ex 8
	fmt.Println("Hands-on ex 8, 9, 10")
	SolHO08lvl4()
}
