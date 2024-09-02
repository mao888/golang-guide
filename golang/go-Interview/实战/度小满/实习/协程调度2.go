/**
    @author: edy
    @since: 2022/9/20
    @desc: //TODO
**/
package main

import (
	"fmt"
	"time"
)

func g1(p chan int) {
	for i := 1; i <= 10; i++ {
		p <- i
		if i%2 == 1 {
			fmt.Println("g1", i)
		}
	}
}

func g2(p chan int) {
	for i := 1; i <= 10; i++ {
		<-p
		if i%2 == 0 {
			fmt.Println("g2", i)
		}
	}
}

func main() {
	msg := make(chan int)
	go g1(msg)
	go g2(msg)

	time.Sleep(time.Second * 1)
}
