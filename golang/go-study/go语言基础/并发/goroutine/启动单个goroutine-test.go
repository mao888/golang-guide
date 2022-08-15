package main

import (
	"fmt"
	"time"
)

func hello3() {
	fmt.Println("hello")
}

func main() {
	go hello3() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second)
}
