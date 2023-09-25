// Q26_fibonacci_go.go
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	term := 25
	i := 0
	c := make(chan int)
	start := time.Now()

	go fibnterms(term, c)
	for {
		if result, ok := <-c; ok {
			fmt.Printf("fibonacci(%d) is: %d\n", i, result)
			i++
		} else {
			end := time.Now()
			delta := end.Sub(start)
			fmt.Printf("longCalculation took this amount of time: %s\n", delta)
			os.Exit(0) // os.Exit(0) 用于确保主函数（main）在所有斐波那契数列的项都计算和打印出来后立即结束 	0 通常表示程序正常退出，而非零的值表示有错误或异常情况。
		}
	}
}

func fibnterms(term int, c chan int) {
	for i := 0; i <= term; i++ {
		c <- fibonacci(i)
	}
	close(c)
}
