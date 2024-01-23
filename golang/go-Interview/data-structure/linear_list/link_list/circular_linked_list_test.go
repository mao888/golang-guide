package linkedlist

import (
	"fmt"
	"github.com/mao888/mao-gutils/constants"
	"testing"
)

func TestCircularLinkedList(t *testing.T) {
	// New 创建一个具有n个元素的环形链表
	ring := New(5)

	// Len 返回环形链表中的元素个数，复杂度O(n)
	len := ring.Len()
	fmt.Println("环形链表中的元素个数:", len) // 5

	// 给循环链表赋值
	for i := 0; i < len; i++ {
		ring.Value = i
		ring = ring.Next()
	}

	// Next 返回后一个元素，r不能为空
	for j := 0; j < len; j++ {
		fmt.Println(ring.Value)
		ring = ring.Next() // 0 1 2 3 4
	}

	// Prev 返回前一个元素，r不能为空。
	for j := 0; j < len; j++ {
		ring = ring.Prev() // 4 3 2 1 0
		fmt.Println("Prev", ring.Value)
	}

	// Do 对链表的每一个元素都执行f（正向顺序），注意如果f改变了*r，Do的行为是未定义的
	ring.Do(func(a any) {
		fmt.Println("Do", a.(int)) // 0 1 2 3 4
	})

	// 返回移动n个位置（n>=0向前移动，n<0向后移动）后的元素，r不能为空。
	// 将指针向前移动三步
	ring = ring.Move(constants.NumberThree)
	ring.Do(func(a any) {
		fmt.Println("Move", a.(int)) // 3 4 0 1 2
	})

	// Link连接r和s，并返回r原本的后继元素r.Next()。r不能为空。
	// 如果r和s指向同一个环形链表，则会删除掉r和s之间的元素，删掉的元素构成一个子链表，返回指向该子链表的指针（r的原后继元素）；
	// 如果没有删除元素，则仍然返回r的原后继元素，而不是nil。
	// 如果r和s指向不同的链表，将创建一个单独的链表，将s指向的链表插入r后面，返回s原最后一个元素后面的元素（即r的原后继元素）。
	ring2 := New(5)
	len2 := ring2.Len()
	for i := 0; i < len2; i++ {
		ring2.Value = 6
		ring2 = ring2.next
	}
	rs := ring.Link(ring2)
	rs.Do(func(a any) {
		fmt.Println("Link:", a.(int)) // 4 0 1 2 3 6 6 6 6 6
	})

	// 删除链表中n % r.Len()个元素，从r.Next()开始删除。
	// 如果n % r.Len() == 0，不修改r。返回删除的元素构成的链表，r不能为空。
	rs.Unlink(3)
	rs.Do(func(a any) {
		fmt.Println("Unlink", a.(int)) // 4 3 6 6 6 6 6
	})
}
