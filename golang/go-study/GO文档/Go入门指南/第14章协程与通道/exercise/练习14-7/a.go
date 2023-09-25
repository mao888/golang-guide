package main

import "fmt"

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i // 发送数字到通道
	}
	close(ch) // 关闭通道，这很重要，否则 main 函数中的 for 循环会永远阻塞等待
}

func main() {
	// 使用 Go 协程和通道
	ch := make(chan int) // 创建一个整数通道
	done := make(chan bool)

	go tel(ch) // 在一个新的 Go 协程中启动 tel 函数

	// 从通道中获取并打印数字
	//for i := range ch {
	//	fmt.Printf("The counter is at %d\n", i)
	//}				 // solution 1
	go cus(ch, done) // solution 2
	<-done
}

func cus(ch chan int, done chan bool) {
	// 匿名函数
	go func() {
		for i := range ch {
			fmt.Printf("The counter is at %d\n", i)
		}
		done <- true
	}()
}
