/**
    @author:huchao
    @data:2022/3/8
    @note:剑指 Offer II 068. 查找插入位置
**/
package main

import "fmt"

func main()  {
	arr := []int{1,3,5,6}
	fmt.Println(searchInsert(arr,7))
}

func searchInsert(nums []int, target int) int {
	return binarySearch(nums,0,len(nums)-1,target)
}

// 二分查找算法(查找的数唯一)
/**
 * @Author huchao
 * @Description //TODO 二分查找算法(查找的数唯一)
 * @Date 19:57 2022/3/8
 * @Param arr   数组
 * @Param left   左边的索引
 * @Param right  右边的索引
 * @Param findVal 要查找的值
 * @return  如果找到就返回下标，如果没有找到，就返回 -1
 **/
func binarySearch(arr []int,left int,right int,findVal int) int {

	// 当 left > right 时，说明递归整个数组，但是没有找到
	for left <= right {

		mid := (right + left) / 2
		midVal := arr[mid]

		if findVal > midVal { // 向右移动
			return binarySearch(arr, mid+1, right, findVal)
		} else if findVal < midVal { // 向左移动
			return binarySearch(arr, left, mid-1, findVal)
		} else {
			return mid
		}
	}
	return left

}