/**
    @author:Hasee
    @data:2022/3/18
    @note:
**/
package main

import "fmt"

func main()  {
	slice := []int{1,2,3,4,5,6}
	RangeSlice(slice)
	fmt.Println(slice)
}

func RangeSlice(slice []int) {
	for index, value := range slice {
		_, _ = index, value
	}
}