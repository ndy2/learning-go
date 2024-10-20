package main

import (
	"errors"
	"fmt"
)

func doSomething() error {
	return errors.New("original error")
}

func main() {
	err := doSomething()
	if err != nil {
		wrappedErr := fmt.Errorf("failed to do something: %w", err) // %w로 원래 오류를 래핑
		fmt.Println(wrappedErr)                                     // 새로운 메시지와 원래 오류 출력

		if orgErr := errors.Unwrap(wrappedErr); orgErr != nil { // Unwrap으로 원래 오류를 가져옴
			fmt.Println(orgErr) //
		}
	}
}
