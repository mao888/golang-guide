package main

import "fmt"

func main() {
	// 分糖果
	fmt.Println(giveT(10))
}

func answer(x int) int {
	// 分糖果
	a := giveT(x)

	// 二
	b := NineToTen(a)
	// 三
	c := Joseph(b)
	return c
}

// 一、分糖果
func giveT(x int) int {
	a1 := x
	a2 := a1 + 5
	a3 := a2 + 5
	a4 := a3 + 5
	a5 := a4 + 5
	return a5
}

// 二
func NineToTen(n int) []int {
	num := make([]int, 0)
	for n > 0 {
		temp := n % 9
		num = append(num, temp)
		n = n / 9
	}
	for i := 0; i < len(num)/2; i++ {
		num[i], num[len(num)-i-1] = num[len(num)-i-1], num[i]
	}
	return num
}

// 三
type Child struct {
	next *Child
	val  int
}

func NewChild(val int) *Child {
	return &Child{
		next: nil,
		val:  val,
	}
}

func Detection(value, counts int) int {
	head := NewChild(1)
	current := head
	for i := 2; i < value; i++ {
		newNode := NewChild(i)
		current.next = newNode
		current = newNode
	}
	// 循环链表
	current.next = head

	count := 0

	// 尾节点
	pre := current
	// 首节点
	current = current.next

	for current.next != current {
		count++
		if count == counts {
			// 删除这个节点
			pre.next = current.next
			count = 0
		}
		pre = current
		current = current.next
	}
	return current.val
}

func Joseph(num []int) int {
	res := Detection(len(num), 5)
	return res
}
