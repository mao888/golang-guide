/**
    @author:Hasee
    @data:2022/3/20
    @note:
**/
package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2) //设置计数器，数值即为goroutine的个数
	go func() {
		//Do some work
		time.Sleep(1*time.Second)

		fmt.Println("Goroutine 1 finished!")
		wg.Done() //goroutine执行结束后将计数器减1
	}()

	go func() {
		//Do some work
		time.Sleep(2*time.Second)

		fmt.Println("Goroutine 2 finished!")
		wg.Done() //goroutine执行结束后将计数器减1
	}()

	wg.Wait() //主goroutine阻塞等待计数器变为0
	fmt.Printf("All Goroutine finished!")
}