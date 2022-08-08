/**
    @author:Hasee
    @data:2022/3/13
    @note:
**/
package main

import "fmt"

func printArray(array *[3]int) {
	for i := range array {
		fmt.Println(array[i])
	}
}

func deferFuncParameter2() {
	var aArray = [3]int{1, 2, 3}

	defer printArray(&aArray)

	aArray[0] = 10 	// 修改数组第一个元素
	return
}
// 延迟函数printArray()的参数在defer语句出现时就已经确定了，
// 即数组的地址，
// 由于延迟函数执行时机是在return语句之前，
// 所以对数组的最终修改值会被打印出来。

func main() {
	deferFuncParameter2()
}
