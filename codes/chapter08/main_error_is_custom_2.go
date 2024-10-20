package main

import (
	"errors"
	"fmt"
	"os"
)

type MyErr struct {
	Err string
}

func (me MyErr) Error() string {
	return me.Err
}

func someFunc() error {
	return MyErr{Err: "some error"}
}

var ErrMyErr = MyErr{Err: "My Special Error"}

func otherFunc() error {
	return ErrMyErr
}

func (me MyErr) Is(target error) bool {
	return false
}

func main() {
	//err := someFunc()
	err := otherFunc()

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
			//} else if errors.Is(err, MyErr{Err: "some error"}) {
		} else if errors.Is(err, ErrMyErr) {
			fmt.Println("That's my error")
		} else {
			fmt.Println("Unknown error")
		}
	}
}
