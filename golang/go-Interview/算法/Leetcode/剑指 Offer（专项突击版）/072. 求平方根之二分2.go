/**
    @author:huchao
    @data:2022/3/10
    @note:剑指 Offer II 072. 求平方根之二分2
**/
package main

import (
	"fmt"
)

func main()  {
	fmt.Println(mySqrt3(8))
}

func mySqrt3(x int) int {
	var left,right = 0 ,x
	for left < right {
		mid := left + (right-left)/2
		if x / mid > mid {
			left = mid + 1
		}else if x/mid < mid {
			right = mid
		}else {
			return mid
		}
	}
	if left * left > x {
		return left - 1
	}else {
		return left
	}
}