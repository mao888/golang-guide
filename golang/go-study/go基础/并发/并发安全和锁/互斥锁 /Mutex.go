/**
    @author: huchao
    @since: 2022/8/2
    @desc: //互斥锁
**/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // 等待组
var x int64
var mutex sync.Mutex

// add 对全局变量x执行5000次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		mutex.Lock() // 修改x前加锁
		x += 1
		mutex.Unlock() // 改完解锁
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
