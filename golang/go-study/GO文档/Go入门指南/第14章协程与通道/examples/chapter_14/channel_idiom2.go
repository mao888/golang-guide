// 给通道使用 for 循环
package main

import (
	"fmt"
	"time"
)

func main() {
	suck4(pump4())
	time.Sleep(1e9)
}

func pump4() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck4(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
