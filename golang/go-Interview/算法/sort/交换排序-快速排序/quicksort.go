package main

import (
	"fmt"
)

/**
 * @Author huChao
 * @Description 交换排序-快速排序
 * @Date 20:56 2022/12/21
 * @Param 1. left 表示 数组左边的下标
 * @Param 2. right 表示数组右边的下标
 * @Param 3. array 表示要排序的数组
 * @return []int
 **/

func QuickSort(left int, right int, array []int) []int {
	l := left
	r := right
	// pivot 是中轴， 支点
	pivot := array[(left+right)/2]
	//for 循环的目标是将比 pivot 小的数放到 左边
	//  比 pivot 大的数放到 右边
	for l < r {
		//从  pivot 的左边找到大于等于pivot的值
		for array[l] < pivot {
			l++
		}
		//从  pivot 的右边边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		// 1 >= r 表明本次分解任务完成, break
		if l >= r {
			break
		}
		//交换
		array[l], array[r] = array[r], array[l]
		//优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	// 如果  1== r, 再移动下
	if l == r {
		l++
		r--
	}
	// 向左递归
	if left < r {
		QuickSort(left, r, array)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, array)
	}
	return array
}

func main() {
	arr := []int{-9, 78, 0, 23, -567, 70, 123, 90, -23}
	quickSort := QuickSort(0, len(arr)-1, arr)
	fmt.Println(quickSort) // [-567 -23 -9 0 23 70 78 90 123]
}
