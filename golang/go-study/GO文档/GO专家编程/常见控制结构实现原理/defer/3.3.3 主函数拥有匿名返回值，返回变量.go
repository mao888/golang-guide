/**
    @author:Hasee
    @data:2022/3/13
    @note:
**/
package main

import "fmt"

func main()  {
	fmt.Println(foo())
}

func foo() int {
	var i int

	defer func() {
		i++
	}()

	return i
}