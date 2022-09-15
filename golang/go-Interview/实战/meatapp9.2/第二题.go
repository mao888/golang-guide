package main

import "fmt"

func main() {
	//answer2()
	//var arr [][]int = [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	var arr []int = []int{1, 4, 56, 1, 6, 2}
	fmt.Println(product(arr))
}

//func answer2(x []int) int {
//	// 第一问
//	a := anser(x)
//	// 第二问
//	b := product(a)
//	// 第三问
//	c := indexForValue(b)
//	return c
//}

// 一、数组排序
func anser(x []int) []int {
	// 定义新的数组
	var new []int
	// 先将数组排序
	c := sort(x)
	// 再把每个值依次取出
	for _, value := range c {
		for i := 0; i < 4; i++ {
			new = append(new, value)
		}
	}
	return new
}

func sort(x []int) []int {
	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i] > x[j] {
				temp := x[i]
				x[i] = x[j]
				x[j] = temp
			}
		}
	}
	return x
}

// 第二问：
func product(a []int) []int {
	length := len(a)

	// L 和 R 分别表示左右两侧的 和 列表
	L, R, answer1 := make([]int, length), make([]int, length), make([]int, length)
	// L[i]为索引 i 左侧所有元素的和
	// 0 左侧没有元素，所以 L[0] = 1
	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = a[i-1] + L[i-1]
	}

	// R[i]为索引 i 左侧所有元素的和
	// length -1 右侧没有元素，所以 R[length - 1] = 1
	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = a[i+1] + R[i+1]
	}

	// 对于索引 i, 除 a[i] 之外其余各元素的和就是左侧所有元素的和 加 右侧所有元素的和
	for i := 0; i < length; i++ {
		answer1[i] = (L[i] + R[i]) * 3
	}
	return answer1
}

func Three(num []int) []int {
	ints := make([]int, len(num))
	count := 0
	for i := 0; i < len(num); i++ {
		count += num[i]
	}
	for i := 0; i < len(num); i++ {
		ints[i] = 3 * (count - num[i])
	}
	return ints
}

// 三、
func indexForValue(num []int) int {
	n := len(num)
	if n == 0 || n == 1 {
		return num[n]
	} else if n >= 2 && n < 5 {
		return num[n/16-1]
	} else if n >= 5 && n < 10 {
		return num[n/8-1]
	} else if n >= 10 && n < 21 {
		return num[n/4-1]
	} else if n >= 21 && n < 42 {
		return num[n/2-1]
	} else {
		return num[42]
	}
}
