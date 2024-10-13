package main

import "fmt"

func main() {
	// 배열 선언 (기본 값)
	var a1 [3]int
	fmt.Println(a1) // [0 0 0]

	// 배열 리터럴
	var a2 = [3]int{1, 2, 3}
	fmt.Println(a2) // [1 2 3]

	// 희소 배열 (sparse array)
	var a3 = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(a3) // [1 0 0 0 0 4 6 0 0 0 100 15]

	var a4 = [...]int{1, 2, 3}
	fmt.Println(a4) // [1 2 3]

	// 배열 비교
	fmt.Println(a2 == a4) // true

	// 다차원배열
	var x [2][3]int
	fmt.Println(x) // [[0 0 0] [0 0 0]]

	x[0][1] = 1
	fmt.Println(x[0][1]) // 1
	fmt.Println(x)       // [[0 1 0] [0 0 0]]

	// 배열의 길이 (len)
	fmt.Println(len(x))    //2
	fmt.Println(len(x[0])) //3
}
