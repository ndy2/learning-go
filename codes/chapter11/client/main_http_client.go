package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"https://jsonplaceholder.typicode.com/todos/1",
		nil,
	)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status code: %d", res.StatusCode))
	}
	fmt.Println(res.Header)                     // type: http.Header
	fmt.Println(res.Header.Get("Content-Type")) // type: string
	fmt.Println(res.Body)                       // type: *http.body

	// NOTE
	// - res.Body is an io.ReadCloser
	//   - 한 번 읽으면 다시 읽을 수 없어!
	//   - 아래 코드를 실행하고 나면 json decoding 시 EOF 에러가 발생한다.
	//bytes, _ := io.ReadAll(res.Body)
	//fmt.Println(string(bytes)) // string

	var data struct {
		ID        int    `json:"id"`
		Content   string `json:"title"`
		Completed bool   `json:"completed"`
		UserID    int    `json:"userId"`
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	// `%+v`는 Go의 `fmt` 패키지에서 제공하는 포맷팅 옵션으로, 구조체의 필드 이름과 값을 함께 출력한다.
	fmt.Printf("%+v\n", data)
}
