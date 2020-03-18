package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Hello World", "test")
	_, _ = fmt.Println(n, err)
}
