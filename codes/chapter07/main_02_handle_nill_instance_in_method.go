package main

import "fmt"

type name struct {
	firstName string
	lastName  string
}

func (n *name) rename(newFirstName, newLastName string) {
	// without this check
	// panic: runtime error: invalid memory address or nil pointer dereference
	// on the line: nilName.rename("n", "dy")
	if n == nil {
		return
	}

	n.firstName = newFirstName
	n.lastName = newLastName
}

func main() {
	ndy := name{"n", "dy"}
	ndy.rename("nam", "deukyun")

	fmt.Println(ndy) // {nam deukyun}

	nilName := (*name)(nil)
	// equivalent to:
	//var nilName *name
	nilName.rename("n", "dy")

	fmt.Println(nilName) // <nil>
}
