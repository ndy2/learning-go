package main

import "fmt"

type Adder struct {
	x int
}

func (a Adder) Add(y int) int {
	return a.x + y
}

func main() {
	adder := Adder{x: 10}
	fmt.Println(adder.Add(5)) // 15

	// method value
	var f1 func(y int) int = adder.Add
	fmt.Println(f1(5)) // 15

	// method expression
	// java 의 method reference 와 비슷한 개념
	var f2 func(Adder, int) int = Adder.Add
	fmt.Println(f2(adder, 5)) // 15
}
