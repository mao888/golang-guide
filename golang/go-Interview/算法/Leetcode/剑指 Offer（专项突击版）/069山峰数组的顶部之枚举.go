/**
    @author:huchao
    @data:2022/3/9
    @note:剑指 Offer II 069. 山峰数组的顶部
**/
package main

import (
	"fmt"
)

func main()  {
	arr := []int{1,3,5,4,2}
	fmt.Println(peakIndexInMountainArray2(arr))
}

func peakIndexInMountainArray2(arr []int) int {
	ans := 1
	for i := 2; i < len(arr)-1; i++ {
		if arr[ans] > arr[i] {
			return ans
		}
		ans = i
	}
	return ans
}