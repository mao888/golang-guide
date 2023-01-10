package linkedlist

import (
	"fmt"
)

// 线性表 - 双向链表
// https://geekr.dev/posts/go-data-structure-linked-list

// 定义节点
type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

// 添加节点
func addNode(t *Node, v int) int {
	if head == nil {
		t = &Node{v, nil, nil}
		head = t
		return 0
	}

	if v == t.Value {
		fmt.Println("节点已存在:", v)
		return -1
	}

	// 如果当前节点下一个节点为空
	if t.Next == nil {
		// 与单链表不同的是每个节点还要维护前驱节点指针
		temp := t
		t.Next = &Node{v, temp, nil}
		return -2
	}

	// 如果当前节点下一个节点不为空
	return addNode(t.Next, v)
}

// 遍历链表
func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> 空链表!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}

	fmt.Println()
}

// 反向遍历链表
func reverse(t *Node) {
	if t == nil {
		fmt.Println("-> 空链表!")
		return
	}

	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}

	for temp.Previous != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}

	fmt.Printf("%d -> ", temp.Value)
	fmt.Println()
}

// 获取链表长度
func size(t *Node) int {
	if t == nil {
		fmt.Println("-> 空链表!")
		return 0
	}

	n := 0
	for t != nil {
		n++
		t = t.Next
	}

	return n
}

// 查找节点
func lookupNode(t *Node, v int) bool {
	if head == nil {
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

// 初始化头节点
var head = new(Node)

func main() {
	fmt.Println(head)
	head = nil
	// 遍历链表
	traverse(head)
	// 新增节点
	addNode(head, 1)
	// 再次遍历
	traverse(head)
	// 继续添加节点
	addNode(head, 10)
	addNode(head, 5)
	addNode(head, 100)
	// 再次遍历
	traverse(head)
	// 添加已存在节点
	addNode(head, 100)
	fmt.Println("链表长度:", size(head))
	// 再次遍历
	traverse(head)
	// 反向遍历
	reverse(head)

	// 查找已存在节点
	if lookupNode(head, 5) {
		fmt.Println("该节点已存在!")
	} else {
		fmt.Println("该节点不存在!")
	}
}
