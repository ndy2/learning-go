package main

import (
	"errors"
	"fmt"
	"os"
)

// calcRemainderAndMod 함수는 두 정수를 나눈 결과와 나머지를 반환한다.
// 0으로 나눈 경우 예외를 발생시킨다.
func calcRemainderAndMod(a int, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return a / b, a % b, nil
}

func main() {
	fmt.Println(calcRemainderAndMod(10, 3))
	fmt.Println(calcRemainderAndMod(10, 0))

	// 예외 처리
	a := 10
	b := 0
	if result, remainder, err := calcRemainderAndMod(a, b); err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println(result, remainder)
	}

	// 예외 처리
	result, remainder, err := calcRemainderAndMod(a, b)
	if err != nil {
		fmt.Println("Error occurred:", err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}
