package main

import (
	"fmt"
	"sync"
)

// 多生产者场景下 channel 关闭的问题
func main() {
	ch := make(chan int)
	workN := 5 // 生产者数
	var wg sync.WaitGroup
	wg.Add(workN)

	// 1、启动多个 Goroutine，通过闭包将 wg 传入，并且在完成任务后调用 wg.Done()。
	for i := 0; i < workN; i++ {
		go func(i int) {
			n := i * i
			ch <- n
			wg.Done()
		}(i)
	}

	// close channel 启动另一个 Goroutine 用来关闭 channel，这个 Goroutine 会阻塞，直到所有的任务完成，最后关闭 channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 3、不停地从 channel 中读取数据，直到 channel 关闭，然后结束程序
	for i := range ch {
		fmt.Println(i)
	}
}
