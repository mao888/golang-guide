package main

import "fmt"

func tel2(ch chan int, done chan struct{}) {
	for i := 0; i < 15; i++ {
		ch <- i // 发送数字到通道
	}
	close(ch)          // 关闭数据通道
	done <- struct{}{} // 发送完成信号
}

func main() {
	ch := make(chan int)        // 创建一个整数通道
	done := make(chan struct{}) // 创建一个结束信号通道

	go tel2(ch, done) // 在一个新的 Go 协程中启动 tel 函数

	// 使用 select 语句来从 ch 和 done 中接收数据
	for {
		select {
		case i, ok := <-ch:
			if ok {
				fmt.Printf("The counter is at %d\n", i)
			}
		case <-done:
			fmt.Println("Finished receiving!")
			return
		}
	}
}
