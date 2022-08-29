package main

import "fmt"

func TestSlice(arr []int) {
	arr = append(arr, 3)
	arr = append(arr, 4)
}

func main() {
	arr := make([]int, 1, 10)
	arr = append(arr, 1)
	arr = append(arr, 2)
	TestSlice(arr)
	fmt.Printf("%d,%d,%d", arr[0], arr[1], len(arr)) // 0,1,3
}
