/**
    @author: huchao
    @since: 2022/8/3
    @desc: //TODO 有缓冲channel
**/
package main

import "fmt"

func producer2(out chan<- int) {
	for i := 0; i < 10; i++ {
		data := i * i
		fmt.Println("生产者生产数据:", data)
		out <- data // 缓冲区写入数据
	}
	close(out) //写完关闭管道
}

func consumer2(in <-chan int) {

	// 无需同步机制，先做后做
	// 没有数据就阻塞等
	for data := range in {
		fmt.Println("消费者得到数据：", data)
	}

}

func main() {
	// 传参的时候显式类型像隐式类型转换，双向管道向单向管道转换
	ch := make(chan int, 5) // 添加缓冲区，5

	go producer2(ch) // 子go程作为生产者
	consumer2(ch)    // 主go程作为消费者
}
