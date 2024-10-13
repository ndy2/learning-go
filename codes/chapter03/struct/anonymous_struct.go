package main

import "fmt"

func main() {
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "John"
	person.age = 30
	person.pet = "dog"
	fmt.Printf("%T\n", person) // struct { name string; age int; pet string }
	fmt.Println(person)        // {John 30 dog}

	pet := struct {
		name string
		kind string
	}{
		name: "Buddy",
		kind: "dog",
	}
	fmt.Printf("%T\n", pet) // struct { name string; kind string }
	fmt.Println(pet)        // {Buddy dog}
}
