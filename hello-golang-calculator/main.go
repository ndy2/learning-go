package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a string: ")
	input, err := reader.ReadString('\n') // 줄바꿈까지 입력 받음
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	var a, b int
	var operator string

	_, err = fmt.Sscanf(input, "%d %s %d", &a, &operator, &b)
	if err != nil {
		fmt.Println("Invalid input. Please enter in the format: number operator number (e.g., 3 + 4)")
		return
	}

	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b != 0 {
			result = a / b
		} else {
			fmt.Println("Cannot divide by zero")
			return
		}
	default:
		fmt.Println("Invalid operator. Please use one of +, -, *, /")
		return
	}

	fmt.Printf("Result: %d\n", result)
}
