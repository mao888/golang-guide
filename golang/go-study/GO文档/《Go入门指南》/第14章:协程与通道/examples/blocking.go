// blocking.go
// throw: all goroutines are asleep - deadlock!
package main

import (
	"fmt"
	"sync"
)

func f1(in chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("in:===", <-in)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	out := make(chan int)
	//out := make(chan int, 1) // solution 2
	//go f1(out, &wg) // solution 1
	out <- 2
	go f1(out, &wg) // solution 2
	wg.Wait()
}

// 这段代码会造成死锁。问题的根源是在 out <- 2 处。在此处，主 goroutine 尝试向 out 通道发送一个整数值，但是没有其他的 goroutine 在接收这个值，导致主 goroutine 被阻塞。
//这是一个常见的 Go 并发模型中的问题。当一个 goroutine 尝试向一个无缓冲的通道发送数据时，它会阻塞，直到另一个 goroutine 从该通道读取数据。但在这段代码中，在发送数据之前没有启动任何接收数据的 goroutine，所以会出现死锁。
// 有两种解决方案：
// 向通道发送数据之前，启动一个 goroutine 来接收数据。
// 将通道的容量设置为 1，这样就可以发送一个数据，然后再阻塞。
