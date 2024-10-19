package main

import "fmt"

type MailCategory int

const (
	UncaughtException MailCategory = iota
	Personal
	Spam
	Social
	Advertisement
)

func main() {
	fmt.Println(UncaughtException) // 0
	fmt.Println(Personal)          // 1
	fmt.Println(Spam)              // 2
	fmt.Println(Social)            // 3
	fmt.Println(Advertisement)     // 4

	fmt.Println(MailCategory(0))  // UncaughtException
	fmt.Println(MailCategory(1))  // Personal
	fmt.Println(MailCategory(2))  // Spam
	fmt.Println(MailCategory(3))  // Social
	fmt.Println(MailCategory(4))  // Advertisement
	fmt.Println(MailCategory(5))  // 5
	fmt.Println(MailCategory(-1)) // -1

	fmt.Println(MailCategory(0) == UncaughtException) // true
	fmt.Println(MailCategory(1) == Personal)          // true
	fmt.Println(MailCategory(2) == Spam)              // true
	fmt.Println(MailCategory(3) == Social)            // true
	fmt.Println(MailCategory(4) == Advertisement)     // true
	fmt.Println(MailCategory(5) == 5)                 // true
	fmt.Println(MailCategory(-1) == -1)               // true

	fmt.Printf("%T\n", UncaughtException) // main.MailCategory
}
