package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func Test4(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg := sync.WaitGroup{}
	wg.Add(1) //注意这里为什么是1不是2：虽然是两个携程同时在运行，但是wg.Wait()只需要知道一个goroutine执行完毕，即不等待。后续通过context来终止携程运行。
	c1 := make(chan int, 1)
	defer close(c1)
	c2 := make(chan int, 1)
	defer close(c2)

	c1 <- 0
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context, c1, c2 chan int) {
		for {
			select {
			case index := <-c1:
				if index%2 == 0 {
					fmt.Println("1--", index)
					fmt.Println("=====================goroutine1 print :", s[index])
					if index+1 >= len(s) {
						wg.Done()
					} else {
						c2 <- index + 1
					}
				}
			case <-ctx.Done():
				fmt.Println("----------------------------goroutine2 exit ...")
				return
			}
		}
	}(ctx, c1, c2)

	go func(ctx context.Context, c1, c2 chan int) {
		for {
			select {
			case index := <-c2:
				if index%2 == 1 {
					fmt.Println("2--", index)
					fmt.Println("=====================goroutine1 print :", s[index])
					if index+1 >= len(s) {
						wg.Done()
					} else {
						c1 <- findex + 1
					}
				}
			case <-ctx.Done():
				fmt.Println("----------------------------goroutine2 exit ...")
				return
			}
		}
	}(ctx, c1, c2)

	wg.Wait()
	cancel()
}
