package main

import (
	"fmt"
)

func InsertSort(arr []int) []int {

	//完成第一次，给第二个元素找到合适的位置并插入
	for i := 1; i < len(arr); i++ {

		insertVal := arr[i]
		insertIndex := i - 1 // 下标

		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次插入后 %v\n", i, arr)
	}
	return arr
}
func main() {
	arr := []int{23, 0, 12, 56, 34, -1, 55}

	insertSort := InsertSort(arr)
	fmt.Println(insertSort)
}
