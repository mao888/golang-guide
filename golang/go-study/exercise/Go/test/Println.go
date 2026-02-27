package main

import "fmt"

func main() {
	//n, e := fmt.Println()
	arr := []int{5, 2, 9, 1, 5, 6}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// 快速排序
func quickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition(left, right, arr)
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
	}
}

func partition(left, right int, arr []int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}
