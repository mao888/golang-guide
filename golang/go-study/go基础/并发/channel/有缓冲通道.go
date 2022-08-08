package main

import "fmt"

func rec2(c chan int) (int int, err error) {
	ret := <-c
	return ret, nil
}

func main() {
	ch := make(chan int, 2)
	for i := 0; i < cap(ch); i++ {
		ch <- i
		a, err := rec2(ch)
		if err != nil {
			return
		}
		fmt.Println(a)
	}
}
