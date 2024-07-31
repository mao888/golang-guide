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

// 有一个两层或者多层的二叉树，按照广度优先遍历，第一层从左到右，第二层从右到左。/**
// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int

	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		leveSize := len(queue)
		var currentLevel []int

		for i := 0; i < leveSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// 偶数 从右到左
			if level%2 == 0 {
				currentLevel = append(currentLevel, node.Val)
			} else { // 奇数 从左到右,直接将节点加入当前层的列表
				currentLevel = append(currentLevel, node.Val)
			}

		}
	}

	return result
}
func main() {
	fmt.Println("hello")
}
