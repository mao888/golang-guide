/**
    @author:huchao
    @data:2022/3/9
    @note:剑指 Offer II 068. 查找插入位置 二分2
**/
package main

import "fmt"

func main()  {
	arr := []int{1,3,5,6}
	fmt.Println(searchInsert2(arr,7))
}

func searchInsert2(nums []int, target int) int {
	return binarySearch2(nums,0,len(nums)-1,target)
}

func binarySearch2(arr []int,left int,right int,findVal int) int {

	// 当 left > right 时，说明递归整个数组，但是没有找到
	for left <=right {

		mid := (right + left) / 2
		midVal := arr[mid]

		if findVal > midVal { // 向右移动
			left = mid+1
		} else if findVal < midVal { // 向左移动
			right = mid-1
		} else {
			return mid
		}
	}
	return left

}
