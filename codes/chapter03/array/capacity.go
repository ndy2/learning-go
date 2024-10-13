package main

import "fmt"

func main() {
	// capacity 가 1024 보다 작은 경우에는 2배씩 확장하고
	// 그보다 큰 경우에는 25% 씩 확장한다.
	var x []int
	fmt.Println(x, len(x), cap(x)) // [] 0 0
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x)) // [10] 1 1
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x)) // [10 20] 2 2
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x)) // [10 20 30] 3 4
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x)) // [10 20 30 40] 4 4
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x)) // [10 20 30 40 50] 5 8

	// make 는 타입, 길이, 그리고 선택적으로 cap 를 지정하여 슬라이스를 만들 수 있다.
	s1 := make([]int, 5)
	fmt.Println(s1, len(s1), cap(s1)) // [0 0 0 0 0] 5 5

	// make 는 타입, 길이, 그리고 선택적으로 cap 를 지정하여 슬라이스를 만들 수 있다.
	s2 := make([]int, 5, 20)
	fmt.Println(s2, len(s2), cap(s2)) // [0 0 0 0 0] 5 20

	// 슬라이스의 길이가 0일 때 그 슬라이스가 반드시 nil 이 되는 것은 아닙니다.
	// - 슬라이스는 make 나 []int{}로 생성되었을 경우 nil 이 아니며, 메모리가 할당된 상태입니다.
	// - 슬라이스가 초기화되지 않았을 때 nil 일 수 있습니다.
	fmt.Println(make([]int, 0) == nil) // false
}
