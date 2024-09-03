package main

import "fmt"

// 题目：给一个无序的数组，找出第k大的数

// partation 对数组进行划分，返回划分后的索引,小于等于pivot的在左边，大于pivot的在右边
func partation(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

// 递归的在子数组中选择第k大的元素
func quickSelect(nums []int, left, right, k int) int {
	if left <= right {
		pivotIndex := partation(nums, left, right)
		if pivotIndex == k { // pivot 索引等于k，返回该元素
			return nums[pivotIndex]
		} else if pivotIndex < k { // pivot 索引小于k，说明k在右边，继续递归
			return quickSelect(nums, pivotIndex+1, right, k)
		} else { // pivot 索引大于k，说明k在左边，继续递归
			return quickSelect(nums, left, pivotIndex-1, k)
		}
	}
	return -1
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))
}
