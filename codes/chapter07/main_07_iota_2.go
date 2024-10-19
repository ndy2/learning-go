package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota) // 1024
	MB                    // 1048576
	GB                    // 1073741824
)

const (
	A = iota // 0
	_        // 1 (건너뜀)
	C        // 2
)

func main() {
	fmt.Println(KB) // 1024
	fmt.Println(MB) // 1048576
	fmt.Println(GB) // 1073741824

	fmt.Println(A) // 0
	fmt.Println(C) // 2
}
