// gofibonacci_select.go
package main

import "fmt"

func goFibonacciSelect(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	goFibonacciSelect(c, quit)
}

/* Output:
1
1
2
3
5
8
13
21
34
55
quit
*/
