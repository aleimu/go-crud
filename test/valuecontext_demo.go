package main

import (
	"context"
	"fmt"
	"time"
)

func HandelRequest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running, parameter: ", ctx.Value("parameter"))
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx0, _ := context.WithTimeout(context.Background(), time.Second*4)
	ctx := context.WithValue(ctx0, "parameter", "1") //context支持嵌套，来完成各种组合技

	go HandelRequest(ctx)

	time.Sleep(10 * time.Second)
}
