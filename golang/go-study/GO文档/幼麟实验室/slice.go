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

	// 测试扩容
	testGrowSlice()
}

func testGrowSlice() {
	s := make([]int, 0)

	oldCap := cap(s)

	for i := 0; i < 2048; i++ {
		s = append(s, i)

		newCap := cap(s)

		if newCap != oldCap {
			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
			oldCap = newCap
		}
	}
}

/*
[0 ->   -1] cap = 0     |  after append 0     cap = 1		// 0 * 2
[0 ->    0] cap = 1     |  after append 1     cap = 2		// 1 * 2
[0 ->    1] cap = 2     |  after append 2     cap = 4		// 2 * 2
[0 ->    3] cap = 4     |  after append 4     cap = 8		// 4 * 2
[0 ->    7] cap = 8     |  after append 8     cap = 16		// 8 * 2
[0 ->   15] cap = 16    |  after append 16    cap = 32		// 16 * 2
[0 ->   31] cap = 32    |  after append 32    cap = 64		// 32 * 2
[0 ->   63] cap = 64    |  after append 64    cap = 128		// 64 * 2
[0 ->  127] cap = 128   |  after append 128   cap = 256		// 128 * 2
[0 ->  255] cap = 256   |  after append 256   cap = 512		// 256 * 2
[0 ->  511] cap = 512   |  after append 512   cap = 848		// 512 * 1.66
[0 ->  847] cap = 848   |  after append 848   cap = 1280	// 848 * 1.5
[0 -> 1279] cap = 1280  |  after append 1280  cap = 1792	// 1280 * 1.4
[0 -> 1791] cap = 1792  |  after append 1792  cap = 2560	// 1792 * 1.42
*/
