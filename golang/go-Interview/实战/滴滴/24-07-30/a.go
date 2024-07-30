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
