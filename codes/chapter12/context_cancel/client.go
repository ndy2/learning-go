package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var client = http.Client{}

func callBoth(ctx context.Context, errVal string, slowURL string, fastURL string) {
	// Create a new context with a cancel function.
	// The cancel function is called when the function returns.
	// NOTE
	// - 취소 가능한 컨텍스트를 만들 때마다 취소 함수를 호출 해야 한다!
	// - `defer cancel()`를 사용하여 최종적으로 호출 되게 하자.
	// - 처음 호출 이후의 수행은 무시 되기 때문에 여러 번 불려도 상관 없다.
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// WaitGroup을 사용하여 두 개의 고루틴이 모두 완료될 때까지 기다린다.
	var wg sync.WaitGroup
	// 두 개의 고루틴을 시작한다.
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := callServer(ctx, "slow", slowURL)
		// 에러가 발생하면 취소 함수를 호출한다.
		// Go 에서 `context`가 취소되면, 해당 `context`의 **취소 신호**가 모든 관련된 Goroutine 에 전달됩니다.
		if err != nil {
			cancel()
		}
	}()
	go func() {
		defer wg.Done()
		err := callServer(ctx, "fast", fastURL+"?error="+errVal)
		// 에러가 발생하면 취소 함수를 호출한다.
		// Go 에서 `context`가 취소되면, 해당 `context`의 **취소 신호**가 모든 관련된 Goroutine 에 전달됩니다.
		if err != nil {
			cancel()
		}
	}()

	// 두 개의 고루틴이 모두 완료될 때까지 기다린다.
	wg.Wait()
	fmt.Println("done with both")
}

// callServer 는 주어진 URL 로 HTTP GET 요청을 보내고 결과를 출력한다.
func callServer(ctx context.Context, label string, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(label, "request err:", err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(label, "response err:", err)
		return err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(label, "read err:", err)
		return err
	}
	result := string(data)
	if result != "" {
		fmt.Println(label, "result:", result)
	}
	if result == "error" {
		fmt.Println("cancelling from", label)
		return errors.New("error happened")
	}
	return nil
}
