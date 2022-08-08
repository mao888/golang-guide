/**
    @author:huchao
    @data:2022/3/10
    @note:剑指 Offer II 070. 排序数组中只出现一次的数字
**/
package main

import "fmt"

func main()  {
	arr := []int{3,3,7,7,10,11,11}
	fmt.Println(singleNonDuplicate(arr))
}

/**
 * 二分查找过程中中间值nums[mid]的特征？有三种情况
 * 	1、它跟它后面一个数字相同
 * 	2、它跟它前面一个数字相同
 * 	3、它就是答案
 *  如果mid所对应的一对数字下标是(奇数，偶数)，那么目标一定在mid之前，如果下标是(偶数，奇数)，目标一定在mid之后
 */
func singleNonDuplicate(nums []int) int {
	var left, right = 0,len(nums)-1
	for left <= right {
		mid := left + (right - left) / 2
		if mid < len(nums) && nums[mid] == nums[mid+1] {  // mid 在左
			if mid % 2 == 0 {	//偶 奇 左指针右移
				left = mid + 2
			} else {	//奇 偶 右指针左移
				right = mid -1
			}
		}else if nums[mid] == nums[mid-1] {		// mid 在右
			if mid % 2 == 0 {	//奇 偶 右指针左移
				right = mid -2
			}else {		//偶 奇 左指针右移
				left = mid + 1
			}
		}else {
			return nums[mid]
		}
	}
	return 0
}