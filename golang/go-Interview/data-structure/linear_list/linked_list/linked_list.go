package linkedlist

import (
	"fmt"
	"strings"
)

// ListInterface 数据结构之线性表--单链表
type ListInterface interface {
	NewListNode(val int) *ListNode
	InsertNode(n0 *ListNode, P *ListNode)
	RemoveNode(n0 *ListNode)
	Access(head *ListNode, index int) *ListNode
	FindNode(head *ListNode, target int) int
	PrintLinkedList(node *ListNode)
}

/*ListNode 链表结点结构体 */
type ListNode struct {
	Val  interface{} // 结点值
	Next *ListNode   // 指向下一结点的指针（引用）
}

// NewListNode 构造函数，创建一个新的链表
func (l *ListNode) NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

/*InsertNode 在链表的结点 n0 之后插入结点 P */
func (l *ListNode) InsertNode(n0 *ListNode, P *ListNode) {
	n1 := n0.Next
	n0.Next = P
	P.Next = n1
}

/*RemoveNode 删除链表的结点 n0 之后的首个结点 */
func (l *ListNode) RemoveNode(n0 *ListNode) {
	if n0.Next == nil {
		return
	}
	// n0 -> P -> n1
	P := n0.Next
	n1 := P.Next
	n0.Next = n1
}

/*Access 访问链表中索引为 index 的结点 */
func (l *ListNode) Access(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		head = head.Next
		if head == nil {
			return nil
		}
	}
	return head
}

/*FindNode 在链表中查找值为 target 的首个结点 */
func (l *ListNode) FindNode(head *ListNode, target int) int {
	index := 0
	for head != nil {
		if head.Val == target {
			return index
		}
		head = head.Next
		index++
	}
	return -1
}

//PrintLinkedList PrintLinkedList Print a linked list
func (l *ListNode) PrintLinkedList(node *ListNode) {
	if node == nil {
		return
	}
	var builder strings.Builder
	for node.Next != nil {
		builder.WriteString(fmt.Sprintf("%v", node.Val) + " -> ") // 使用 fmt.Sprintf 将interface value转换为字符串
		node = node.Next
	}
	builder.WriteString(fmt.Sprintf("%v", node.Val))
	fmt.Println(builder.String())
}
