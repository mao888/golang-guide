package main

import (
	"fmt"
	"time"
)

// 用两个携程实现一个从1到10到交替顺序打印

var flagChan = make(chan int)

func work1() {
	for i := 1; i <= 10; i++ {
		flagChan <- 1 // 塞入
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}

func work2() {
	for i := 1; i <= 10; i++ {
		_ = <-flagChan // 	取出
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	go work1();
	go work2();

	time.Sleep(3 * time.Second)
}