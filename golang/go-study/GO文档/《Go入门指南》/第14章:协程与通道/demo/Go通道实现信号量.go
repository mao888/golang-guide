package main

import (
	"fmt"
	"sync"
	"time"
)

// 模拟某种工作，需要限制并发访问
func doWork(id int, sem chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	sem <- true // 请求访问

	fmt.Printf("Worker %d started\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done\n", id)

	<-sem // 释放访问
}

func main() {
	var wg sync.WaitGroup
	sem := make(chan bool, 3) // 信号量，同时只允许3个协程访问资源

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go doWork(i, sem, &wg)
	}

	wg.Wait()
}

//在上面的代码中：
//
//- 我们定义了一个模拟工作的 doWork 函数，该函数使用了信号量（通过通道实现）来控制访问。
//
//- 我们创建了一个带缓冲的通道 sem 作为信号量。它的容量是 3，这意味着同时只允许 3 个协程访问资源。
//
//- 在 doWork 函数内部，我们首先发送一个值到 sem 通道来请求访问。如果 sem 通道的容量已满，则此操作将阻塞，直到有其他协程从该通道中读取一个值为止。
//
//- 完成工作后，我们从 sem 通道读取一个值，以释放一个访问权限，从而允许其他等待中的协程进入。
//
//- 此代码展示了如何使用 Go 的通道来实现信号量模式，从而限制并发访问某个资源。
