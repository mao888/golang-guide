/**
    @author:Hasee
    @data:2022/3/19
    @note:
**/
package main

import (
	"fmt"
	"runtime"
)

func main()  {
	a := runtime.GOMAXPROCS(10)
	fmt.Println(a)
}
