package main

import (
	"fmt"
	"github.com/ndy2/learning-go/package-example/formatter"
	"github.com/ndy2/learning-go/package-example/math"
)

func main() {
	fmt.Println("Hello, Package!")
	num := math.Double(2)
	fmt.Println(num) // 4
	output := print.Format(num)
	fmt.Println(output) // num: 4
}
