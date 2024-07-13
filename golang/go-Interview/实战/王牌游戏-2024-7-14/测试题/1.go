package main

import (
	"fmt"
)

// TreeNode 表示二叉树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PathInfo 存储路径信息
type PathInfo struct {
	length int         // 路径长度
	sum    int         // 路径节点值之和
	nodes  []*TreeNode // 路径上的节点
}

// DFS 遍历树并收集路径信息
func DFS(node *TreeNode, length int, sum int, path []*TreeNode, paths *[]PathInfo) {
	if node == nil {
		return
	}

	// 将当前节点加入路径
	path = append(path, node)
	length++
	sum += node.Val

	// 添加当前路径信息到 paths
	pathCopy := make([]*TreeNode, len(path))
	copy(pathCopy, path)
	*paths = append(*paths, PathInfo{length, sum, pathCopy})

	// 继续遍历左右子树
	DFS(node.Left, length, sum, path, paths)
	DFS(node.Right, length, sum, path, paths)
}

// hasCommonNode 检查两个路径是否有共同节点
func hasCommonNode(path1, path2 []*TreeNode) bool {
	nodeSet := make(map[*TreeNode]struct{})
	for _, node := range path1 {
		nodeSet[node] = struct{}{}
	}
	for _, node := range path2 {
		if _, exists := nodeSet[node]; exists {
			return true
		}
	}
	return false
}

// FindAllPairs 查找所有符合条件的节点对
func FindAllPairs(root *TreeNode) [][2]*TreeNode {
	var paths []PathInfo
	DFS(root, 0, 0, []*TreeNode{}, &paths)

	var result [][2]*TreeNode
	n := len(paths)

	// 查找路径长度相同且路径和相同的节点对
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if paths[i].length == paths[j].length && paths[i].sum == paths[j].sum {
				if !hasCommonNode(paths[i].nodes, paths[j].nodes) {
					result = append(result, [2]*TreeNode{paths[i].nodes[0], paths[j].nodes[0]})
				}
			}
		}
	}

	return result
}

func main() {
	// 构建一个示例二叉树
	//        1
	//       / \
	//      2   3
	//     /|   |\
	//    4 5   6 7
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	// 查找符合条件的节点对
	pairs := FindAllPairs(root)
	for _, pair := range pairs {
		fmt.Printf("节点对: (%d, %d)\n", pair[0].Val, pair[1].Val)
	}
}
