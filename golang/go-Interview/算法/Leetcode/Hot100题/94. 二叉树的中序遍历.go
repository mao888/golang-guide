/**
    @author:Hasee
    @data:2022/6/27
    @note:
**/
package main

import "fmt"

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var inorder func(node *TreeNode)
	var res []int
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}

func main() {
	root := &TreeNode{Val: 1}
	null := &TreeNode{Val: 0}
	root.Left = null
	r := &TreeNode{Val: 2}
	rL := &TreeNode{Val: 3}
	root.Right = r
	r.Left = rL
	//var root = [2]int{1, nil}
	num := inorderTraversal(root)
	println(len(num))
	for i := 0; i < len(num); i++ {
		fmt.Print(num[i])
	}
}
