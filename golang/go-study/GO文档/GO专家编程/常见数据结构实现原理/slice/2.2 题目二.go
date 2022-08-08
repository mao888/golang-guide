/**
    @author:huchao
    @data:2022/3/10
    @note:2.2 题目二
**/
package main

import (
	"fmt"
)

//函数AddElement()接受一个切片和一个元素，把元素append进切片中，并返回切片
func AddElement(slice []int, e int) []int {
	return append(slice, e)
}

func main() {
	var slice []int
	fmt.Println(cap(slice))
	slice = append(slice, 1,2,3)
	fmt.Println(cap(slice))

	newSlice := AddElement(slice, 4)		// 向切片append进第4个元素同时定义一个新的切片newSlice
	//newSlice = AddElement(newSlice,5)

	//fmt.Println(cap(newSlice))
	fmt.Printf("%p,%p",slice,newSlice)
	fmt.Println(&slice[0] == &newSlice[0])	// 新切片newSlice与旧切片slice是否共用一块存储空间

	//append函数执行时会判断切片容量是否能够存放新增元素，
	//如果不能，则会重新申请存储空间，新存储空间将是原来的2倍或1.25倍（取决于扩展原空间大小），
	//本例中实际执行了两次append操作，第一次空间增长到4，
	//所以第二次append不会再扩容，所以新旧两个切片将共用一块存储空间。程序会输出”true”。
}