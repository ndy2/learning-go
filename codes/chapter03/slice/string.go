package main

import "fmt"

func main() {
	var s string = "안녕 golang"
	fmt.Println(s)        // 안녕 golang
	fmt.Printf("%T\n", s) // string
	fmt.Println(len(s))   // 13
	fmt.Println(s[0:1])   // �
	fmt.Println(s[0:2])   // �
	fmt.Println(s[0:3])   // 안
}
