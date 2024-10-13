package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4) // x: [1 2 3 4] cap: 5
	y := x[:2:2]              // y: [1 2] cap: 2
	z := x[2:4:4]             // z: [3 4] cap: 2

	fmt.Println(cap(x), cap(y), cap(z)) // 5 2 2
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("x:", x) // x: [1 2 3 4 60]
	fmt.Println("y:", y) // y: [1 2 30 40 50]
	fmt.Println("z:", z) // z: [3 4 70]

	y[0] = 100
	fmt.Println("x:", x)
}
