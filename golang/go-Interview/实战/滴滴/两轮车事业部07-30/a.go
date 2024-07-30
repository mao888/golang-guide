package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//1. 给出一个长度非空的整形数组，元素值先递增后递减，找出数组中的最大值
//eg: []int{1, 2, 3, 4, 6, 3, 1} 输出6
//eg: []int{1, 2, 3, 4, 6} 输出6
//eg: []int{6, 4, 3, 2, 1} 输出6

// 初始化两个指针 left 和 right 分别指向数组的第一个元素和最后一个元素。
// 进入循环，直到 left 指针小于 right 指针：
// a. 计算中间指针 mid。
// b. 如果 mid 指向的元素大于其右侧的元素，说明最大值可能在 mid 或其左侧，因此将 right 指针移动到 mid。
// c. 否则，说明最大值可能在 mid 的右侧（但不在 mid 本身，因为数组是先递增后递减的），因此将 left 指针移动到 mid + 1。
// 当循环结束时，left 和 right 指针会相遇，它们指向的元素就是数组中的最大值。
// 这个算法的时间复杂度是 O(logn)，因为它每次都将搜索范围减半，直到找到最大值。
func findMaxValue(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left)/2

		// 如果 mid 指向的元素 > 右侧的，说明 max 在 mid 或其左侧
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			// max 在 mid 右侧
			left = mid + 1
		}
	}
	// left == right 此时 max
	return arr[left]
}

func main() {

	fmt.Println(findMaxValue([]int{1, 2, 3, 4, 6, 3, 1}))
	fmt.Println(findMaxValue([]int{1, 2, 3, 4, 6, 7}))
	fmt.Println(findMaxValue([]int{7, 6, 4, 3, 2, 1}))
}
