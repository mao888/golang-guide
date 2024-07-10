package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func f() {
	defer fmt.Println("D")
	fmt.Println("F")
}

func main() {
	defer fmt.Println("N")
	f()
	fmt.Println("M")
}
