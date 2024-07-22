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

// 入参是一个数组，里面的数字不重复。返回的结果是一个数组的数组。入参 [1 2 3] 返回入参的全排列[[1 2 3][1 3 2][2 1 3][2 3 1][3 1 2][3 2 1]]
// 全排列
func permutation(nums []int) [][]int {
	var (
		result [][]int
		// 递归函数，接收遍历到的索引
		backtrack func(int)
	)
	backtrack = func(first int) {
		// 所有都数都填完了，找到了一个排列，填入result
		if first == len(nums) {
			// 复制当前的排列
			result = append(result, append([]int(nil), nums...))
		}
		// 交换，递归，交换
		for i := first; i < len(nums); i++ {
			nums[first], nums[i] = nums[i], nums[first]
			// 递归
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrack(0)
	return result
}
func main() {
	// 入参是一个数组，里面的数字不重复。返回的结果是一个数组的数组。
	// 入参 [1 2 3] 返回入参的全排列[[1 2 3][1 3 2][2 1 3][2 3 1][3 1 2][3 2 1]]
	arr := []int{1, 2, 3}
	p := permutation(arr)
	fmt.Println(p) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 2 1] [3 1 2]]
}
