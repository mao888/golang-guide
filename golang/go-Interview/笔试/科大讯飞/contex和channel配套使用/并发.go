/**
    @author: HuChao
    @since: 2022/8/4
    @desc: //TODO
**/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	result := make(chan int, 0)
	asyncDoStuffWithTimeout(ctx, result)
	fmt.Printf("restult get: %v", <-result)
}

func asyncDoStuffWithTimeout(ctx context.Context, result chan int) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("ctx is done, %v", ctx.Err())
			result <- 0
			return
		case <-time.After(2 * time.Second):
			fmt.Println("set result")
			result <- 10
		}
	}()
}
