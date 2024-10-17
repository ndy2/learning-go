package main

import "fmt"

type Person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p Person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func main() {
	i := 10
	s := "Hello"
	p := Person{age: 20, name: "Alice"}
	modifyFails(i, s, p)

	// i, s, and p remain unchanged
	fmt.Println(i, s, p)
}
