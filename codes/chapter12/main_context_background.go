package main

import (
	"context"
	"fmt"
)

func logic(ctx context.Context, info string) (string, error) {
	select {
	case <-ctx.Done():
		fmt.Println("logic done")
		return "", ctx.Err()
	default:
		fmt.Println("logic start")
		return info, nil
	}
}

func main() {
	ctx := context.Background()
	result, err := logic(ctx, "hello")
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("result: ", result)
	}
}
