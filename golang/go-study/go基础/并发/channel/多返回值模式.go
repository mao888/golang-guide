package main

import "fmt"

func f2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

//	for range接收值
func f3(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch) // 当向通道中发送完数据时，我们可以通过close函数来关闭通道
	f2(ch)
}
