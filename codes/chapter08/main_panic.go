package main

import (
	"fmt"
)

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v) // prints "runtime error: integer divide by zero"
		}
	}()
	fmt.Println(60 / i)
}

func main() {
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}
