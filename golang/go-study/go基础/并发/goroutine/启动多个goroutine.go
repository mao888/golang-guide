package main

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

func hello2(i int) {
	defer wg2.Done() // goroutine结束就登记-1
	fmt.Println("hello", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg2.Add(1) //	启动一个goroutine就登记+1
		go hello2(i)
	}
	wg2.Wait() // 等待所有登记的goroutine都结束
}
