package main

import (
	"fmt"
	"sync"
)

//	两个goroutine交替打印1-100之间的奇数和偶数

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			ch <- 1 // 塞入
			//奇数
			if i%2 == 1 {
				fmt.Println("线程1打印:", i)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			<-ch // 取出
			//偶数
			if i%2 == 0 {
				fmt.Println("线程2打印:", i)
			}
		}
	}()
	wg.Wait()
}
