/**
    @author:huchao
    @data:2022/3/10
    @note: 2.1 题目一
**/
package main

import "fmt"

func main() {
	var array [10]int		// 定义了一个10个长度的整型数组array

	var slice = array[5:6]	// 定义了一个切片slice，切取数组的第6个元素

	fmt.Println(array)		// [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(slice)		// [0]
	fmt.Println("lenth of slice: ", len(slice))			// 1
	fmt.Println("capacity of slice: ", cap(slice))		// 5
	fmt.Println(&slice[0] == &array[5])					// true
	// slice根据数组array创建，与数组共享存储空间，
	// slice起始位置是array[5]，长度为1，容量为5，slice[0]和array[5]地址相同。
}