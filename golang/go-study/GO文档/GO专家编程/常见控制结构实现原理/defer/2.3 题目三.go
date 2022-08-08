/**
    @author:Hasee
    @data:2022/3/13
    @note:
**/
package main

import "fmt"

func main()  {
	fmt.Println(deferFuncReturn3())
}

func deferFuncReturn3() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i
}