/**
    @author:huchao
    @data:2022/3/12
    @note: 3.3 使用数组创建Slice
**/
package main

import "fmt"

func main()  {

	//使用数组来创建Slice时，Slice将与原数组共用一部分内存。
	var array [10]int
	slice := array[5:7]
	fmt.Println(len(slice))	// 2  切片从数组array[5]开始，到数组array[7]结束（不含array[7]），即切片长度为2
	fmt.Println(cap(slice))	// 5  数组后面的内容都作为切片的预留内存，即capacity为5。
	fmt.Println(&array[5] == &slice[0]) // true

	fmt.Printf("%p,%p",slice,array)
}