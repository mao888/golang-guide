package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// 定义链表的节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func getNthToLast(head *ListNode, n int) *ListNode {
	// 快慢指针，指向头节点
	slow, fast := head, head

	// 快指针先向前移动 n 步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 快慢指针同时向前移动，直到快指针到达链表的末尾
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 此时慢指针所在的位置就是倒数第 n 个节点
	return slow

}
func main() {
	// 创建一个简单的链表： 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	// 获取倒数第二个节点
	node := getNthToLast(head, 2) // 输出：4

	// 打印倒数第二个节点的值
	fmt.Println(node.Val) // 输出：4
}
