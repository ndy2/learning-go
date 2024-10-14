package main

import "fmt"

func main() {
outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outerLoop // outerLoop 레이블이 붙은 바깥쪽 for 문을 종료
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("Done")
}
