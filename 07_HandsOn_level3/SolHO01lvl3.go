package main

import "fmt"

// SolHO01lvl3 : Solution ex 1 lvl 3
func SolHO01lvl3() {
	s := ""
	for i := 1; i <= 100; i++ {
		s = fmt.Sprintf("%v\t%v ", s, i) // String concatenation
		if i%10 == 0 {
			fmt.Println(s)
			s = ""
		}
	}
}

// SolHO02lvl3 : Solution ex 2 lvl 3
func SolHO02lvl3() {
	fmt.Println("Sol ex 2 lvl 3")
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

// SolHO03lvl3 : Sol ex 3
func SolHO03lvl3() {
	bd := 1985
	for bd <= 2020 {
		fmt.Println(bd)
		bd++
	}
}

// SolHO04lvl3 : Sol ex 4
func SolHO04lvl3() {
	bd := 1985
	for {
		if bd > 2020 {
			break
		}
		fmt.Println(bd)
		bd++
	}

}

// SolHO05lvl3 : Sol ex 5
func SolHO05lvl3() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("The remainder of %v divided by 4 is %v\n", i, i%4)
	}
}

// SolHO06lvl3 : Sol ex 6
func SolHO06lvl3() {
	q := "print"

	if q == "print" {
		fmt.Println(q)
	}
}

// SolHO07lvl3 : Sol ex 7
func SolHO07lvl3() {
	q := "prxnt"

	if q == "prnt" {
		fmt.Println("print")
	} else if q == "do not print" {
		fmt.Println("do not print")
	} else {
		fmt.Println("what?")
	}
}

// SolHO08lvl3 : Sol ex 8
func SolHO08lvl3() {
	a := 1
	b := 2
	switch {
	case (a == b):
		fmt.Println("a = b")
	default:
		fmt.Println("a != b")
	}
}

// SolHO09lvl3 : Sol ex 9
func SolHO09lvl3() {
	favSport := "t"
	switch favSport {
	case "r":
		fmt.Println("r")
	case "t":
		fmt.Println("t")
	default:
	}
}
