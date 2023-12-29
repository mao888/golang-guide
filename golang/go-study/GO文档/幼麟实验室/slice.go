package main

import (
	"fmt"
)

func main() {
	var ints []int = make([]int, 2, 5)
	ints = append(ints, 1)
	ints[0] = 1
	//ints[3] = 2 // panic: runtime error: index out of range [3] with length 2

	// new
	ps := new([]string)
	//(*ps)[0] = "eggo" // panic: runtime error: index out of range [0] with length 0
	*ps = append(*ps, "hello")

	fmt.Println(*ps) // [hello]
	fmt.Println(&ps) // 0xc0000b4000
	fmt.Println(ps)  // &[hello]

	// 数组
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s1 []int = arr[1:4]
	fmt.Println(s1)               // [1 2 3]
	fmt.Println(len(s1), cap(s1)) // 3 9
	//fmt.Println(s1[3])            // panic: runtime error: index out of range [3] with length 3
	var s2 []int = arr[7:]
	fmt.Println(s2)               // [7 8 9]
	fmt.Println(len(s2), cap(s2)) // 3 3

	// int 扩容
	ints2 := []int{1, 2}                     // len=2, cap=2
	ints2 = append(ints2, 3, 4, 5)           // len=5, cap=6
	ints2 = append(ints2, 3)                 // len=3, cap=4
	ints2 = append(ints2, 4)                 // len=4, cap=4
	ints2 = append(ints2, 5, 6, 7, 8, 9, 10) // len=9, cap=10

	// string 扩容
	a := []string{"a", "b", "c"}
	a = append(a, "eggo")
}
