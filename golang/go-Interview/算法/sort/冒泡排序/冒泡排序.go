/**
    @author:huchao
    @data:2022/3/6
    @note:go冒泡排序
**/
package main

import "fmt"

func main()  {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(BubbleSort(arr))	// 冒泡排序
	fmt.Println(GetMax(arr))		// 冒泡排序获取最大值
}
//冒泡排序获取最大值
func GetMax(arr []int) int {
	for i :=1;i<len(arr);i++ {
		if arr[i-1]>arr[i] {
			arr[i-1],arr[i]=arr[i],arr[i-1];
		}
	}
	return arr[len(arr)-1]
}

//冒泡排序
func BubbleSort(arr []int) []int {
	for i :=0;i<len(arr);i++ {
		for j:=i+1;j<len(arr);j++ {
			if arr[i] > arr[j] {
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}
	return arr
}