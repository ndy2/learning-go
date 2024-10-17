package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// os.Args[0]는 프로그램 이름을 포함하므로, 최소한 하나의 파일 인자가 필요
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// main 함수가 종료될 때 파일을 자동으로 닫도록 합니다. 리소스 누수를 방지합니다.
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
