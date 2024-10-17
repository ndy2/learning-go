package main

import "fmt"

func stringp(s string) *string {
	return &s // 참조 연산자
}

func main() {

	person := struct {
		name *string // 문자열 포인터 타입
		age  int
	}{
		name: stringp("Alice"),
		age:  30,
	}

	fmt.Println("Name:", *person.name) // 역참조 연산자
	fmt.Println("Age:", person.age)
}
