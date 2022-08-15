package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // 等待组
var x int64

// add 对全局变量x执行5000次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		x += 1
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	//runtime.GOMAXPROCS(1)
	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}

/*
开启了两个 goroutine 分别执行 add 函数，
这两个 goroutine 在访问和修改全局的x变量时就会存在数据竞争，
某个 goroutine 中对全局变量x的修改可能会覆盖掉另一个 goroutine 中的操作，
所以导致最后的结果与预期不符。
*/
