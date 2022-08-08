/**
    @author:huchao
    @data:2022/3/9
    @note:剑指 Offer II 069. 山峰数组的顶部
**/
package main

import "fmt"

func main()  {
	arr := []int{0,10,5,2}
	fmt.Println(peakIndexInMountainArray(arr))
}

func peakIndexInMountainArray(arr []int) int {
	return binarySearch3(arr,1,len(arr)-1)
}

func binarySearch3(arr []int,left int,right int) int {

	ans := 0
	// 当 left > right 时，说明递归整个数组，但是没有找到
	for left <=right {

		mid := (right + left) / 2

		if arr[mid] > arr[mid+1] { // 向左移动
			ans = mid
			right = mid - 1
		} else if arr[mid] < arr[mid+1] { // 右移动
			left = mid+1
		} else {
			return mid
		}

	}
	return ans
}
