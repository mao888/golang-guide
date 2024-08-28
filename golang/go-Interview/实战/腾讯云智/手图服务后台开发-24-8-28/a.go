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

type MaxStack struct {
	stack    []int
	maxStack []int
}

func Constructor() MaxStack {
	return MaxStack{
		stack:    []int{},
		maxStack: []int{},
	}
}

func (s *MaxStack) Push(x int) {
	s.stack = append(s.stack, x)
	if len(s.maxStack) == 0 || x >= s.maxStack[len(s.maxStack)-1] {
		s.maxStack = append(s.maxStack, x)
	}
}

func (s *MaxStack) Pop() int {
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if top == s.maxStack[len(s.maxStack)-1] {
		s.maxStack = s.maxStack[:len(s.maxStack)-1]
	}
	return top
}

func (s *MaxStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MaxStack) PeekMax() int {
	return s.maxStack[len(s.maxStack)-1]
}

func (s *MaxStack) PopMax() int {
	maxVal := s.PeekMax()
	buffer := []int{}

	// Pop elements until we find the max
	for s.Top() != maxVal {
		buffer = append(buffer, s.Pop())
	}

	// Pop the max element
	s.Pop()

	// Push back the elements from the buffer
	for i := len(buffer) - 1; i >= 0; i-- {
		s.Push(buffer[i])
	}

	return maxVal
}

func main() {
	stack := Constructor()
	stack.Push(5)
	stack.Push(1)
	stack.Push(5)
	fmt.Println(stack.Top())     // -> 5
	fmt.Println(stack.PopMax())  // -> 5
	fmt.Println(stack.Top())     // -> 1
	fmt.Println(stack.PeekMax()) // -> 5
	fmt.Println(stack.Pop())     // -> 1
	fmt.Println(stack.Top())     // -> 5
}
