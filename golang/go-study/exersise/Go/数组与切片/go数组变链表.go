package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	next *ListNode
}

func SliceToLinkList(nums []int, head *ListNode) *ListNode {
	node := head
	for _, num := range nums {
		temp := ListNode{Val: num}
		head.next = &temp
		head = &temp
	}
	return node.next
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	node := new(ListNode)

	linkNode := SliceToLinkList(arr, node)
	for {
		if linkNode != nil {
			fmt.Println(linkNode.Val)
			linkNode = linkNode.next
			continue
		}
		break
	}
}
