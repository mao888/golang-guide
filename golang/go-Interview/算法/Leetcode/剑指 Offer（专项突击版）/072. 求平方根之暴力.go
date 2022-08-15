/**
    @author:huchao
    @data:2022/3/10
    @note:剑指 Offer II 072. 求平方根
**/
package main

import "fmt"

func main()  {
	var x = 2
	fmt.Println(mySqrt(x))
}

func mySqrt(x int) int {
	var i = 1
	if x == 0 {
		return 0
	}
	for i <= x {
		if x / i > i {
			i = i + 1
		}else if x / i < i {
			return i -1
		}else {
			return i
		}
	}
	return i
}