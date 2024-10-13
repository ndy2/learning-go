package main

import "fmt"

func main() {
	// valueType 을 bool 로 해서 집합 처럼 활용
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true // 항상 값을 true 로 채워 주어야 한다.
	}

	fmt.Println(len(vals), len(intSet)) // 11 8
	fmt.Println(intSet[5])              // true
	fmt.Println(intSet[500])            // false
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}
