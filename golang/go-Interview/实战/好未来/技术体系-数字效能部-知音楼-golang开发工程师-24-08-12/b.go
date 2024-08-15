package main

import (
	"fmt"
)

//求子数组的最大和
//题目:输入一个整形数组，数组里有正数也有负数。 数组中连续的一个或多个整数组成一个子数组，每个子数组都有一个和。 求所有子数组的和的最大值。要求时间复杂度为 O(n)。
//例如输入的数组为1, -2, 3, 10, -4, 7, 2, -5，和最大的子数组为3, 10, -4, 7, 2， 因此输出为该子数组的和18。

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	localMax := nums[0]
	globalMax := nums[0]

	for i := 1; i < len(nums); i++ {
		// 判断是开始一个新的子数组还是从上一个子数组中继续
		localMax = max(nums[i], localMax+nums[i])
		// 更新全局最大和
		globalMax = max(globalMax, localMax)
	}
	return globalMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, -2, 3, 10, -4, 7, 2, -5}
	fmt.Println("最大子数组和为:", maxSubArray(nums))
}
