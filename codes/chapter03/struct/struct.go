package main

import "fmt"

func main() {

	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person
	bob := person{}

	fmt.Println(fred, bob)   // { 0 } { 0 }
	fmt.Println(fred == bob) // true

	julia := person{
		name: "Julia",
		age:  42,
		pet:  "cat",
	}
	fmt.Println(julia) // {Julia 42 cat}
	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println(beth) // {Beth 30  }
}
