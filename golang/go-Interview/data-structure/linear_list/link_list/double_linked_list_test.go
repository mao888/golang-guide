/**
    @author:HuChao
    @data:2023/1/10
    @note: doubly linked list test
**/
package linkedlist

import (
	"fmt"
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	// Create a new list and put some numbers in it.
	var l List
	l.New()               // New创建一个链表
	e4 := l.PushBack(4)   // PushBack将一个值为v的新元素插入链表的最后一个位置，返回生成的新元素
	e1 := l.PushFront(1)  // PushFront将一个值为v的新元素插入链表的第一个位置，返回生成的新元素。
	l.InsertBefore(3, e4) // InsertBefore将一个值为v的新元素插入到mark前面，并返回生成的新元素。如果mark不是l的元素，l不会被修改。
	l.InsertAfter(2, e1)  // InsertAfter将一个值为v的新元素插入到mark后面，并返回新生成的元素。如果mark不是l的元素，l不会被修改。
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() { // 遍历得：1 2 3 4
		fmt.Println(e.Value)
	}

	len := l.Len()                // Len返回链表中元素的个数，复杂度O(1)。
	fmt.Println("链表中元素的个数:", len) // 4

	front := l.Front()                   // Front返回链表第一个元素或nil。
	fmt.Println("链表第一个元素:", front.Value) // 1

	back := l.Back()                     // Back返回链表最后一个元素或nil。
	fmt.Println("链表最后一个元素:", back.Value) // 4

	l.PushFrontList(&l) // PushFrontList创建链表other的拷贝，并将拷贝的最后一个位置连接到链表l的第一个位置。
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("PushFrontList:", e.Value) // 1 2 3 4  1 2 3 4
	}

	l.PushBackList(&l) // PushBackList创建链表other的拷贝，并将链表l的最后一个位置连接到拷贝的第一个位置。
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("PushBackList:", e.Value) // 1 2 3 4  1 2 3 4  1 2 3 4 1 2 3 4
	}

	l.MoveToFront(e4) // MoveToFront将元素e移动到链表的第一个位置，如果e不是l的元素，l不会被修改。
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("MoveToFront:", e.Value) // 4  1 2 3 4  1 2 3  1 2 3 4 1 2 3 4
	}

	l.MoveBefore(e1, e4) // MoveBefore将元素e移动到mark的前面。如果e或mark不是l的元素，或者e==mark，l不会被修改。
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("MoveBefore:", e.Value) // 1 4  1 2 3 4  2 3  1 2 3 4 1 2 3 4
	}

	l.Remove(e1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("Remove:", e.Value) // 4  1 2 3 4  2 3  1 2 3 4 1 2 3 4
	}

	l.Init() // Init清空链表。
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() { // 遍历得：空
		fmt.Println("Init清空链表", e.Value)
	}

}
