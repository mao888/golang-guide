package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println("ζ₯ζΆζε", ret)
}
func main() {
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("success")
}
