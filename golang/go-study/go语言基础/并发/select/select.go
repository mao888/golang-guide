package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
			fmt.Println("输入i：", i)
		}
	}
}

//第一次循环时 i = 1，select 语句中包含两个 case 分支，此时由于通道中没有值可以接收，所以x := <-ch 这个 case 分支不满足，而ch <- i这个分支可以执行，会把1发送到通道中，结束本次 for 循环；
//第二次 for 循环时，i = 2，由于通道缓冲区已满，所以ch <- i这个分支不满足，而x := <-ch这个分支可以执行，从通道接收值1并赋值给变量 x ，所以会在终端打印出 1；
//后续的 for 循环以此类推会依次打印出3、5、7、9。
