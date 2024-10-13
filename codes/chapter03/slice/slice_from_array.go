package main

import (
	"fmt"
)

func main() {
	x := [4]int{5, 6, 7, 8}     // this is an array
	y := x[:2]                  // this is a slice
	z := x[2:]                  // this is a slice
	x[0] = 10                   // memory 공유
	fmt.Printf("%T %v\n", x, x) // [4]int [10 6 7 8]
	fmt.Printf("%T %v\n", y, y) // []int [10 6]
	fmt.Printf("%T %v\n", z, z) // []int [7 8]
}
