package main

import (
	"fmt"
	"io"
	"os"
)

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func main() {
	//   什么类型可以作为map 的key,给出例子
	//   1. 数值类型
	map1 := map[int]int{1: 1, 2: 2, 3: 3}
	//   2. 字符串类型
	aMap := map[string]int{"a": 1, "b": 2, "c": 3}
	//   3. 指针类型
	uMap := map[*int]int{new(int): 1, new(int): 2, new(int): 3}
	//   8. interface类型
	map8 := map[interface{}]int{1: 1, 2: 2, 3: 3}
	//   9. channel类型
	map9 := map[chan int]int{make(chan int): 1, make(chan int): 2, make(chan int): 3}
	//   10. 结构体类型
	type A struct {
	}
	map10 := map[A]int{A{}: 1, A{}: 2, A{}: 3}
	//   11. 数组类型
	map11 := map[[3]int]int{[3]int{1, 2, 3}: 1, [3]int{2, 3, 4}: 2, [3]int{3, 4, 5}: 3}
	//   13. 接口类型
	map13 := map[io.Reader]int{os.Stdin: 1, os.Stdout: 2, os.Stderr: 3}
	//   14. bool类型
	map14 := map[bool]int{true: 1, false: 2}
	fmt.Println(map1, aMap, uMap, map8, map9, map10, map11, map13, map14)
}
