/**
    @author:huchao
    @data:2022/3/6
    @note: 283. 移动零
**/
package main

import "fmt"

func moveZeroes(nums []int) []int {
	var left int = 0
	var right int = 0
	for right < len(nums) {
		if nums[right] !=0 {
			nums[left],nums[right] = nums[right],nums[left]
			left++
		}
		right++
	}
	return nums
}

func main()  {
	arr := []int{0,1,0,3,12}
	fmt.Println(moveZeroes(arr))
}