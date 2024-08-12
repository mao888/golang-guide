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

// levelOrder 返回二叉树的层序遍历结果
func levelOrder(root *TreeNode) [][]int {
	var result [][]int // 存储最终的层序遍历结果
	if root == nil {   // 如果根节点为空，直接返回空结果
		return result
	}

	queue := []*TreeNode{root} // 初始化队列，并将根节点加入队列

	// 当队列不为空时，继续处理
	for len(queue) > 0 {
		var level []int // 存储当前层的节点值
		n := len(queue) // 当前层的节点数量

		// 遍历当前层的所有节点
		for i := 0; i < n; i++ {
			node := queue[0]                // 取出队列中的第一个节点
			queue = queue[1:]               // 移除已经处理的节点
			level = append(level, node.Val) // 将节点值加入当前层结果中

			// 如果左子节点不为空，加入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			// 如果右子节点不为空，加入队列
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 将当前层的节点值加入最终结果
		result = append(result, level)
	}

	return result // 返回层序遍历的最终结果
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

	fmt.Println("-----------------")

	// 构建一个示例二叉树
	//        1
	//       / \
	//      2   3
	//     /|   |\
	//    4 5   6 7
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 3}
	root2.Left.Left = &TreeNode{Val: 4}
	root2.Left.Right = &TreeNode{Val: 5}
	root2.Right.Left = &TreeNode{Val: 6}
	root2.Right.Right = &TreeNode{Val: 7}
	result2 := levelOrder(root2)
	for _, level := range result2 {
		fmt.Println(level)
	}
}
