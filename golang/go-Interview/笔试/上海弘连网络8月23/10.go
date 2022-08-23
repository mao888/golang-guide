package main

import (
	"fmt"
	"time"
)

// 请您写一段简单的示例代码来表示Go协程+Channel的用法

func Process(ch chan int) {
	//Do some work...
	time.Sleep(time.Second)

	ch <- 1 //管道中写入一个元素表示当前协程已结束
}

func main() {
	channels := make([]chan int, 10) //创建一个10个元素的切片，元素类型为channel

	for i := 0; i < 10; i++ {
		channels[i] = make(chan int) //切片中放入一个channel
		go Process(channels[i])      //启动协程，传一个管道用于通信
	}

	for i, ch := range channels { //遍历切片，等待子协程结束
		<-ch
		fmt.Println("Routine ", i, " quit!")
	}
}
