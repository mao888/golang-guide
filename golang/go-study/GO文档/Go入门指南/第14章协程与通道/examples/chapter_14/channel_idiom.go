// 通道工厂模式 不将通道作为参数传递给协程，而用函数来生成一个通道并返回（工厂角色）；函数内有个匿名函数被协程调用。
// 通道工厂模式是指当我们需要创建大量的通道时，可以使用通道工厂模式来简化代码。
// 通道工厂模式的实现原理是：将通道的创建和通道的使用分离，将通道的创建封装到一个函数中，通过调用该函数来创建通道，然后再将通道传递给需要使用通道的函数。
// 通道工厂模式的代码如下所示。
package main

import (
	"fmt"
	"time"
)

func main() {
	stream := pump3()
	go suck3(stream)
	time.Sleep(1e9)
}

func pump3() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck3(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
