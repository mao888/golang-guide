/**
    @author:Hasee
    @data:2022/3/13
    @note:
**/
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

func main() {
	ch := make(chan int, 1)
	ch <- 1
	//ch <- 2
	close(ch)
	f2(ch)
}
