package main

var c = make(chan int)
var a string

func f() {
	a = "hello"
	c <- 0
}

func main() {
	go f()
	<-c
	print(a) //	hello
}
