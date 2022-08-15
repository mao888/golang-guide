package main

import "fmt"

// 未初始化的通道类型变量其默认零值是nil
var ch1 chan int   // 声明一个传递整型的通道
var ch2 chan bool  // 声明一个传递布尔型的通道
var ch3 chan []int // 声明一个传递int切片的通道

var ch4 = make(chan int)
var ch5 = make(chan bool, 1) // 声明一个缓冲区大小为1的通道

func main() {
	fmt.Println(ch1) //	<nil>
	fmt.Println(ch2) //	<nil>
	fmt.Println(ch3) //	<nil>

	fmt.Println(ch4) //	0xc000064060
	fmt.Println(ch5) //	0xc000100000
}
