package main

import (
	"errors"
	"fmt"
)

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are allowed")
	}
	return i * 2, nil
}

func main() {
	a, i := doubleEven(3)
	fmt.Println(a, i)             // 0 only even numbers are allowed
	fmt.Printf("%T\n", i)         // *errors.errorString
	fmt.Printf("%T\n", i.Error()) // string
}
