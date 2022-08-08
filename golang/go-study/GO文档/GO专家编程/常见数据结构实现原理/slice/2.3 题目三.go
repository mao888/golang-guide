/**
    @author:huchao
    @data:2022/3/10
    @note: 2.3 题目三
**/
package main

import (
	"fmt"
)

func main() {

	orderLen := 5
	order := make([]uint16, 2 * orderLen)	// 定义一个长度为10的切片order

	// pollorder和lockorder分别是对order切片做了order[low:high:max]操作生成的切片
	pollorder := order[:orderLen:orderLen]		 		// 指的是order的前半部分切片
	lockorder := order[orderLen:][:orderLen:orderLen]	// 指的是order的后半部分切片，即原order分成了两段

	a := order[:5:5]
	b := order[5:][:5:5]

	// 分别打印pollorder和lockorder的容量和长度
	fmt.Println("len(pollorder) = ", len(pollorder))	// 5
	fmt.Println("cap(pollorder) = ", cap(pollorder))	// 5
	fmt.Println("len(lockorder) = ", len(lockorder))	// 5
	fmt.Println("cap(lockorder) = ", cap(lockorder))	// 5
}