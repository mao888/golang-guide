//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
// 构建一个示例二叉树
//        1
//       / \
//      2   3
//     /|   |\
//    4 5   6 7
//root := &TreeNode{Val: 1}
//root.Left = &TreeNode{Val: 2}
//root.Right = &TreeNode{Val: 3}
//root.Left.Left = &TreeNode{Val: 4}
//root.Left.Right = &TreeNode{Val: 5}
//root.Right.Left = &TreeNode{Val: 6}
//root.Right.Right = &TreeNode{Val: 7}

package main

import "fmt"

// 快排
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	// 第一个数为基准
	pivot := arr[left]
	for left < right {
		// 从右向左找到 < pivot的数
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left] = arr[right]
		// 从左向右找到 > pivot的数
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[right] = arr[left]
	}
	// 放置基准数，此时left == right
	arr[left] = pivot
	// 递归排序左右
	quickSort(arr[:left])
	quickSort(arr[left+1:])
	return arr
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println("原数组:", arr)
	sortedArr := quickSort(arr)
	fmt.Println("快速排序后的: ", sortedArr)
}
