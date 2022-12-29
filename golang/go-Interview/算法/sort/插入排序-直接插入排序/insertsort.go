package main

import (
	"fmt"
)

// InsertSort 插入排序-直接插入排序
// 平均时间复杂度o(n^2)
// 最好时间复杂度o(n)
// 空间复杂度o(1)
// 稳定
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

/**
第1次插入后 [23 0 12 56 34 -1 55]
第2次插入后 [23 12 0 56 34 -1 55]
第3次插入后 [56 23 12 0 34 -1 55]
第4次插入后 [56 34 23 12 0 -1 55]
第5次插入后 [56 34 23 12 0 -1 55]
第6次插入后 [56 55 34 23 12 0 -1]
[56 55 34 23 12 0 -1]
*/
