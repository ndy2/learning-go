package main

import "fmt"

func main() {
	var pointerToX *int

	// Dereferencing a nil pointer will cause a panic
	fmt.Println(*pointerToX)
}
