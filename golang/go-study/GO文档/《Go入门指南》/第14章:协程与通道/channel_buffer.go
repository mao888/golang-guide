// 给 channel_block3.go 的通道增加缓冲并观察输出有何不同
package main

import "fmt"
import "time"

func main() {
	c := make(chan int, 50)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}

/* Output:
sending 10
sent 10   // prints immediately
no further output, because main() then stops
*/
