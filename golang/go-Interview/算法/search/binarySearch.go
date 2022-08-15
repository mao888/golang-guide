/**
    @author:Huchao
    @data:2022/3/8
    @note: binarySearch二分查找
**/
package main

import (
	"fmt"
)

func main()  {
	arr := []int{1,8,10,89,1000,1000,1000,1000,1234}
	fmt.Println(binarySearch2(arr,0,len(arr)-1,1000))
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
	if left > right {
		return -1
	}
	mid := (right + left)/2
	midVal := arr[mid]

	if findVal > midVal {	// 向右移动
		return binarySearch(arr,mid+1,right,findVal)
	}else if findVal < midVal{	// 向左移动
		return binarySearch(arr,mid-1,right,findVal)
	}else {
		return mid
	}
}

// 二分查找算法(查找的数有重复)
/**
 * @Author huchao
 * @Description //TODO 二分查找算法(查找的数有重复)
 * @Date 19:59 2022/3/8
 * @Param arr   数组
 * @Param left   左边的索引
 * @Param right  右边的索引
 * @Param findVal 要查找的值
 * @return 如果找到就返回下标列表，如果没有找到，就返回 空
 **/
/*
* 思路分析
* 1. 在找到mid 索引值，不要马上返回
* 2. 向mid 索引值的左边扫描，将所有满足 1000， 的元素的下标，加入到集合ArrayList
* 3. 向mid 索引值的右边扫描，将所有满足 1000， 的元素的下标，加入到集合ArrayList
* 4. 将Arraylist返回
*/
func binarySearch2(arr []int,left int,right int,findVal int) []int {
	arr1 := make([]int,0,len(arr))
	// 当 left > right 时，说明递归整个数组，但是没有找到
	if left > right {
		return arr1
	}
	mid := (right + left)/2
	midVal := arr[mid]

	if findVal > midVal {	// 向右移动
		return binarySearch2(arr,mid+1,right,findVal)
	}else if findVal < midVal{	// 向左移动
		return binarySearch2(arr,left,mid-1,findVal)
	}else {
//	* 思路分析
//	* 1. 在找到mid 索引值，不要马上返回
//	* 2. 向mid 索引值的左边扫描，将所有满足 1000， 的元素的下标，加入到集合ArrayList
//	* 3. 向mid 索引值的右边扫描，将所有满足 1000， 的元素的下标，加入到集合ArrayList
//	* 4. 将Arraylist返回

		//2. 向mid 索引值的左边扫描，将所有满足 1000， 的元素的下标，加入到集合ArrayList
		temp := mid - 1
		for temp < len(arr) && arr[temp] == arr[mid] {
			arr1 = append(arr1, temp)
			temp = temp -1
		}
		arr1 = append(arr1, mid)

		//3. 向mid 索引值的右边扫描，将所有满足 1000， 的元素的下标，加入到集合ArrayList
		temp = mid + 1;
		for temp < len(arr) && arr[temp] == arr[mid] {
			arr1 = append(arr1, temp)
			temp = temp + 1
		}
		return arr1
	}
}