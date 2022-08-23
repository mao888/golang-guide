package main

import "fmt"

//面试日期: 2022-08-23 周二
//
//面试时间: 14:00

//给你一棵二叉树的根节点 root ，二叉树中节点的值 互不相同 。另给你一个整数 start 。在第 0 分钟，感染 将会从值为 start 的节点开始爆发。
//
//每分钟，如果节点满足以下全部条件，就会被感染：
//节点此前还没有感染。
//节点与一个已感染节点相邻。
//
//返回感染整棵树需要的分钟数。

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
//
//(1)
//start = 3
//
//1
/// \
//5   3
///   / \
//7   10  6

//(2)
//
//root = [1], start = 1
//
//(3)
//start = 3
//
//1
//\
//3
/// \
//10  6
///
//11
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 创建节点
func CreateBinaryTree(data int) *TreeNode {
	return &TreeNode{data, nil, nil}
}

// 插入节点
func (node *TreeNode) Insert(n *TreeNode, data int) bool {
	cur := n
	for cur != nil {
		if cur.Val < data {
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

func amountOfTime(root *TreeNode, start int) int {
	var st *TreeNode
	parents := map[*TreeNode]*TreeNode{}

	//  dis 求出start的位置和每个点的父节点
	var dfs func(*TreeNode, *TreeNode)
	dfs = func(node *TreeNode, pa *TreeNode) {
		if node == nil {
			return
		}
		if node.Val == start {
			st = node
		}
		parents[node] = pa
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)

	ans := -1
	vis := map[*TreeNode]bool{nil: true, st: true}
	for q := []*TreeNode{st}; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, node := range tmp {
			if node != nil {
				if !vis[node.Left] {
					vis[node.Left] = true
					q = append(q, node.Left)
				}
				if !vis[node.Right] {
					vis[node.Right] = true
					q = append(q, node.Right)
				}
				if p := parents[node]; !vis[p] {
					vis[p] = true
					q = append(q, p)
				}
			}
		}
	}
	return ans
}

func main() {
	var node *TreeNode

	node = CreateBinaryTree(1)
	li := []int{5, 3, 7, 10, 6}

	for _, val := range li {
		node.Insert(node, val)
	}
	a := amountOfTime(node, 3)
	fmt.Println(a)
}
