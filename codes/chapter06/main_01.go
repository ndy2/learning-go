package main

import "fmt"

func main() {
	x := 10
	pointerToX := &x
	fmt.Println("Address of x: ", pointerToX)
	fmt.Println("Value of x: ", *pointerToX)

	z := 5 + *pointerToX
	fmt.Println("Value of z: ", z)
}
