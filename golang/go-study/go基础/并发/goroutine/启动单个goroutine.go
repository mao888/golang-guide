package main

import (
	"fmt"
	"sync"
)

// 当你并不关心并发操作的结
// 或者有其它方式收集并发操作的结果时，
// WaitGroup是实现等待一组并发操作完成的好方法。

// 声明全局等待组变量
var wg sync.WaitGroup

func hello() {
	fmt.Println("hello")
	wg.Done() // 告知当前goroutine完成
}

func main() {
	wg.Add(1) // 登记一个goroutine
	hello()
	fmt.Println("你好")
	wg.Wait() // 阻塞等待登记对goroutine完成
}
