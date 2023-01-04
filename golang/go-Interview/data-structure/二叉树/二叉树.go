package main

import (
	"fmt"
)

// 节点
type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// 创建节点
func CreateBinaryTree(data int) *BinaryTreeNode {
	return &BinaryTreeNode{data, nil, nil}
}

// 插入节点
func (node *BinaryTreeNode) Insert(n *BinaryTreeNode, data int) bool {
	cur := n
	for cur != nil {
		if cur.Data < data {
			if cur.Right != nil {
				cur = cur.Right
			} else {
				cur.Right = CreateBinaryTree(data)
				return true
			}
		} else {
			if cur.Left != nil {
				cur = cur.Left
			} else {
				cur.Left = CreateBinaryTree(data)
				fmt.Println(data, "d---")
				return true
			}
		}
	}
	return false
}

// 层数打印
func (node *BinaryTreeNode) BreadthFirstSearch() []int {
	if node == nil {
		return nil
	}
	var result []int
	par := node
	cur := []*BinaryTreeNode{par}
	for len(cur) > 0 {
		result = append(result, cur[0].Data)
		if cur[0].Left != nil {
			cur = append(cur, cur[0].Left)
		}
		if cur[0].Right != nil {
			cur = append(cur, cur[0].Right)
		}
		cur = cur[1:]
	}
	return result
}

// 前序打印
func (node *BinaryTreeNode) PreOrder(n *BinaryTreeNode) {
	if n != nil {
		fmt.Println(n.Data)
		node.PreOrder(n.Left)
		node.PreOrder(n.Right)
	}
}

// 中序打印
func (node *BinaryTreeNode) InOrder(n *BinaryTreeNode) {
	if n != nil {
		node.InOrder(n.Left)
		fmt.Println(n.Data)
		node.InOrder(n.Right)
	}
}

// 后序打印
func (node *BinaryTreeNode) PostOrder(n *BinaryTreeNode) {
	if n != nil {
		node.InOrder(n.Left)
		node.InOrder(n.Right)
		fmt.Println(n.Data)
	}
}

// 获取树的高度
func (node *BinaryTreeNode) GetHight(n *BinaryTreeNode) int {
	if n == nil {
		return 0
	}
	l := node.GetHight(n.Left)
	r := node.GetHight(n.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}

// 打印叶子节点
func (node *BinaryTreeNode) FindLead(n *BinaryTreeNode) {
	if n != nil {
		if n.Left == nil && n.Right == nil {
			fmt.Println(n.Data)
		}
		node.FindLead(n.Left)
		node.FindLead(n.Right)
	}
}

// 查找指定值的节点
func (node *BinaryTreeNode) FindValueNode(n *BinaryTreeNode, target int) *BinaryTreeNode {
	if n == nil {
		return nil
	} else if n.Data == target {
		return n
	} else {
		cur := node.FindValueNode(n.Left, target)
		if cur != nil {
			return cur
		}
		return node.FindValueNode(n.Right, target)
	}
}

func main() {
	var node *BinaryTreeNode
	// 创建一个根节点
	node = CreateBinaryTree(10)
	li := []int{9, 11, 8, 5, 6, 4, 12, 15, 18, 17}
	// 创建一个二叉树
	for _, val := range li {
		node.Insert(node, val)
	}
	ret := node.BreadthFirstSearch()
	fmt.Println(ret)
	node.PreOrder(node)
	node.InOrder(node)
	node.PostOrder(node)
	res := node.GetHight(node)
	fmt.Println(res)
	node.FindLead(node)
	ref := node.FindValueNode(node, 17)
	fmt.Println(ref)
}
