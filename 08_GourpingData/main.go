package main

import "fmt"

// Grouping Data
func main() {
	fmt.Println("--- Array ---")

	// Array declaration
	var x [5]int // array of 5 elements of type int

	// Assignement
	x[0] = 2
	x[1] = 42
	x[2] = 41
	fmt.Println(x)
	fmt.Println(len(x)) // get the lenght of the array

	// Arrays are not really used. Use slices instead

	fmt.Println("--- Slice ---")
	// y := type{values} --> Composite literal
	// SLICE allows you to group together VALUES of the same TYPE

	y := []int{4, 5, 6, 7, 42}
	fmt.Println(y)
	fmt.Println(y[3]) // Print out the 4th element. Here, 3 is the index

	//  Loop over the value
	fmt.Println("Loop over slice:")
	for i, v := range y {
		fmt.Println(i, v)
	}

	// here, if I do i <= len(y) I got panic.go
	for i := 0; i < len(y); i++ {
		fmt.Println(i, y[i])
	}

	// Slicing a slice
	// [START:END] --> END is not included
	fmt.Println("Slicing a slice:")
	fmt.Println(y[:])
	fmt.Println(y[1:])
	fmt.Println(y[1:3])

	// Append to a slice
	fmt.Println("Append to a slice:")
	y = append(y, 77, 88, 99)
	fmt.Println(y)

	z := []int{123, 456, 789}
	y = append(y, z...) // ... is a variadic parameter and it means take any number of elements for z
	fmt.Println(y)

	// Deleting from a slice
	fmt.Println("Deleting from a slice:")
	y = append(y[:2], y[4:]...)
	fmt.Println(y)

	//  Make
	fmt.Println("Make:")
	// MAKE builds the array underlying a slice
	// make([]T, length, capacity)
	u := make([]int, 10, 12)
	fmt.Println(u)      // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(u)) // 10
	fmt.Println(cap(u)) // 12
	u[0] = 42
	u[9] = 999
	// u[10] = 101 <-- this lead to out of range

	u = append(u, 101)

	fmt.Println(u)      // [42 0 0 0 0 0 0 0 0 999 101]
	fmt.Println(len(u)) // 11
	fmt.Println(cap(u)) // 12

	u = append(u, 102)
	u = append(u, 103)
	u = append(u, 104)

	fmt.Println(u)      // [42 0 0 0 0 0 0 0 0 999 101 102 103 104]
	fmt.Println(len(u)) // 14
	fmt.Println(cap(u)) // 24

	//  Multi-dimensional slice
	fmt.Println("Multi-dimensional slice:")
	// A multi-dimensional slice is like a spreadsheet. You can have a slice of a slice of some type.

	v := []string{"James", "Bond", "007", "Martini"}
	w := []string{"Miss", "Moneypeeny", "Admin", "Peach"}

	vw := [][]string{v, w}
	fmt.Println(vw)
	fmt.Println(vw[1][1])

	fmt.Println("--- Map ---")
	// A map is a unordered key-value store.
	// map[keytype]elementtype{}

	m := map[string]int{
		"key1": 32,
		"key2": 27,
	}

	fmt.Println(m["key2"]) // return 27
	fmt.Println(m["key"])  // return 0 value

	value, ok := m["key"]
	fmt.Printf("value: %v, ok: %v\n", value, ok)

	// good idiomatic way of coding
	// remember: if shortstatement; cond {}
	if value, ok := m["key1"]; ok {
		fmt.Printf("if condition passed. value: %v, ok: %v\n", value, ok)
	}

	// add element & range
	fmt.Println("add element & range")
	m["key3"] = 33
	fmt.Println(m)

	// loop over map
	fmt.Println("loop over map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	// delete
	fmt.Println("delete")
	delete(m, "key3")
	fmt.Println(m)

	delete(m, "key3") // there is no issue to delete no existing key:value
	fmt.Println(m)

	if v, ok := m["key2"]; ok {
		fmt.Println("value:", v, "ok:", ok)
		delete(m, "key2")
	}

	fmt.Println(m)
}
