/**
    @author: huchao
    @data:2022/3/18
    @note:
**/
package main

import (
	"fmt"
	"time"
)

func main() {
	// 程序中声明两个channel，分别为chan1和chan2
	chan1 := make(chan int)
	chan2 := make(chan int)

	//a := new([]string)
	c := "ssdfsads"
	//a1 := make([]string,1,2)
	//*a = append(*a,"1")
	fmt.Println(c[0])
	//(*a)[0] ="1"
	//fmt.Println(*a)

	// 依次启动两个协程，分别向两个channel中写入一个数据就进入睡眠
	go func() {
		chan1 <- 1
		time.Sleep(5 * time.Second)
	}()

	go func() {
		chan2 <- 1
		time.Sleep(5 * time.Second)
	}()

	// select语句两个case分别检测chan1和chan2是否可读，如果都不可读则执行default语句。
	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	default:
		fmt.Println("default")
	}

	fmt.Println("main exit.")

	// select中各个case执行顺序是随机的，如果某个case中的channel已经ready，则执行相应的语句并退出select流程，
	// 如果所有case中的channel都未ready，则执行default中的语句然后退出select流程。

	// 另外，由于启动的协程和select语句并不能保证执行顺序，
	// 所以也有可能select执行时协程还未向channel中写入数据，
	// 所以select直接执行default语句并退出。所以，以下三种输出都有可能：

}