/**
    @author:Hasee
    @data:2022/3/13
    @note:
**/
package main

import "fmt"

func main()  {
	deferFuncParameter()	// 延迟函数fmt.Println(aInt)的参数在defer语句出现时就已经确定了，所以无论后面如何修改aInt变量都不会影响延迟函数。
}

func deferFuncParameter() {
	var aInt = 1

	defer fmt.Println(aInt)

	aInt = 2
	return
}