package main

import "fmt"

type Variable int

const (
	X Variable = iota
	Y
	Z
)

func (v Variable) String() string {
	switch v {
	case X:
		return "X"
	case Y:
		return "Y"
	case Z:
		return "Z"
	}
	return "Unknown"
}

func main() {
	fmt.Println(X.String())
}
