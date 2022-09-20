package main

import (
	"fmt"
	"sync"
)

//整体思路
//简单讲解下，由于goroutine的执行顺序是没有保证的，所以当需要顺序打印时，需要借助其他的一些全局变量控制打印顺序，我此处是通过一个tag，通过tag这个变量控制数据进入chan的顺序，进而使得输出可以按序输出
//需要注意的一个小点
//通过协程进行写入时，不能直接用变量i，需要通过参数把i传递进去，因为i对于这个协程来说是全局的，而i这个全局变量又是在自增的，所以不能直接用i

//使用channel按顺序输出1-10
func orderchan() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	chs := make(chan int)
	//通过一个全局变量控制进channel的顺序
	tag := 1
	for i := 1; i <= 10; i++ {
		go func(value int) {
			//死循环，保证按顺序进chan
			for {
				if tag == value {
					chs <- value
					break
				}
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Print(<-chs)
		wg.Done()
		tag++
	}
	wg.Wait()
}

func main() {
	orderchan()
}
