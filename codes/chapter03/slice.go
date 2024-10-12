package main

import "fmt"

func main() {

	var x []int
	fmt.Println(x)        // []
	fmt.Println(x == nil) // true

	// len
	var s []int
	fmt.Println(len(s))

	// append
	// ERROR - append(s, 10) evaluated but not used
	// Go 는 `CALL BY VALUE`
	// - append 로 전달된 슬라이스는 복사된 값이 함수로 전달된다.
	// - 이 함수는 복사도니 슬라이스에 값들을 추가하고 추가된 복사본을 반환한다.
	// append(s, 10)
	s = append(s, 10)
	fmt.Println(s, len(s)) // [10] 1

	s = append(s, 20, 30, 40)
	fmt.Println(s, len(s)) // [10 20 30 40] 4

	var s2 = []int{1, 2, 3, 4}
	// `...` 연산자 - see 5.1.2 가변 입력 파라미터와 슬라이스
	var s3 = append(s, s2...)
	fmt.Println(s3) // [10 20 30 40 1 2 3 4]
}
