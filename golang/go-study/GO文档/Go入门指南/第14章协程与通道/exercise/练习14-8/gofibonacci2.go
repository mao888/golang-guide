// gofibonacci2.go
package main

import (
	"fmt"
)

func fibonacci3(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci3(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
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
*/
