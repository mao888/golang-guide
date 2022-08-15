/**
    @author: huchao
    @since: 2022/8/3
    @desc: //TODO 无缓冲channel
**/
package main

import "fmt"

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		data := i * i
		fmt.Println("生产者生产数据:", data)
		out <- data // 缓冲区写入数据
	}
	close(out) //写完关闭管道
}

func consumer(in <-chan int) {
	// 同样读取管道
	//for{
	//    val, ok := <- in
	//    if ok {
	//        fmt.Println("消费者拿到数据：", data)
	//    }else{
	//        fmt.Println("无数据")
	//        break
	//    }
	//}

	// 无需同步机制，先做后做
	// 没有数据就阻塞等
	for data := range in {
		fmt.Println("消费者得到数据：", data)
	}

}

func main() {
	// 传参的时候显式类型像隐式类型转换，双向管道向单向管道转换
	ch := make(chan int) //无缓冲channel
	go producer(ch)      // 子go程作为生产者
	consumer(ch)         // 主go程作为消费者
}
