package main

import (
	"fmt"
)

//SelectSort 选择排序-简单选择排序
// 最好、最坏、平均时间复杂度均为：O(n^2),
// 空间复杂度:O(1)
// 稳定性：不稳定 如:5 5 2
func SelectSort(arr []int) []int {
	//1. 先完成将第一个最大值和 arr[0] => 先易后难
	//1 假设  arr[0] 最大值
	for j := 0; j < len(arr)-1; j++ {

		max := arr[j]
		maxIndex := j
		//2. 遍历后面 j + 1---[len(arr) -1] 比较
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] { //找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		fmt.Printf("第%d次 %v\n  ", j+1, arr)
	}
	return arr
}

func main() {
	//定义一个数组
	arr := []int{10, 34, 19, 100, 80, 789}
	selectSort := SelectSort(arr)
	fmt.Println(selectSort)
}
