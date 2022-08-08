package main

import "fmt"

// Producer 返回一个通道
// 并持续将符合条件的数据发送至返回的通道中
// 数据发送完成后会将返回的通道关闭
func Producer() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()
	return ch
}

// Consumer 从通道中接收数据进行计算
func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		fmt.Println("v:", v)
		sum += v
	}
	return sum
}

func main() {
	ch := Producer()
	res := Consumer(ch)
	fmt.Println(res)

	ch2 := Producer2()
	res2 := Consumer2(ch2)
	fmt.Println(res2)

	var ch3 = make(chan int, 1)
	ch3 <- 10
	close(ch3)
	res3 := Consumer2(ch3) // 函数传参时将ch3转为单向通道
	fmt.Println("res3", res3)

	var ch4 = make(chan int, 1)
	ch4 <- 15
	close(ch4)
	var ch5 <-chan int // 声明一个只接收通道ch5
	ch5 = ch4          // 变量赋值时将ch4转为单向通道
	//<-ch5

	res4 := Consumer2(ch5)
	fmt.Println("res4", res4)
}

// Producer2 返回一个接收通道
func Producer2() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// Consumer2 参数为接收通道
func Consumer2(ch <-chan int) int {
	sum := 0
	for v := range ch {
		fmt.Println("v:", v)
		sum += v
	}
	return sum
}
