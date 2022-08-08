/**
    @author:Hasee
    @data:2022/3/13
    @note:
**/
package main

import (
	"fmt"
	"time"
)

func main()  {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
		time.Sleep(time.Millisecond)
	}
	time.Sleep(1*time.Second)
}