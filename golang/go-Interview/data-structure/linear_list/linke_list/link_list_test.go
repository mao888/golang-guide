// File: link_list_test.go
// Created Time: 2022-12-29
// Author: cathay (cathaycchen@gmail.com)

package linkedlist

import (
	"fmt"
	"testing"
)

func TestLikedList(t *testing.T) {
	/* 初始化链表 1 -> 3 -> 2 -> 5 -> 4 */
	// 初始化各个结点
	var li LNode
	n0 := li.NewListNode(1)
	n1 := li.NewListNode(3)
	n2 := li.NewListNode(2)
	n3 := li.NewListNode(5)
	n4 := li.NewListNode(4)

	// 构建引用指向
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	fmt.Println("初始化的链表为")
	li.PrintLinkedList(n0)

	/* 插入结点 */
	li.InsertNode(n0, li.NewListNode(0))
	fmt.Println("插入结点后的链表为")
	li.PrintLinkedList(n0)

	/* 删除结点 */
	li.RemoveNode(n0)
	fmt.Println("删除结点后的链表为")
	li.PrintLinkedList(n0)

	/* 访问结点 */
	node := li.Access(n0, 3)
	fmt.Println("链表中索引 3 处的结点的值 =", node)

	/* 查找结点 */
	index := li.FindNode(n0, 2)
	fmt.Println("链表中值为 2 的结点的索引 =", index)
}
