package main

import "fmt"

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

// 有一个两层或者多层的二叉树，按照广度优先遍历，第一层从左到右，第二层从右到左。/**
// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder 返回二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		var level []int
		n := len(queue)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}

	return result
}

func main() {
	// 构造示例二叉树: [3,9,20,null,null,15,7]
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	// 调用层序遍历函数
	result := levelOrder(root)

	// 打印输出
	for _, level := range result {
		fmt.Println(level)
	}
}
