/*
*

	@author:Hasee
	@data:2022/11/25
	@note:

*
*/
package main

import (
	"fmt"
	"github.com/mao888/mao-gutils/slice"
)

func main() {
	leyangjun1 := []int{1, 3, 5, 6}
	leyangjun2 := []int{1, 3, 5}

	retDiff := DifferenceSet(leyangjun1, leyangjun2)
	s := slice.DifferenceSet(leyangjun1, leyangjun2)
	fmt.Println(retDiff)
	fmt.Println(s)
}

// DifferenceSet 返回两个slice切片的差集
func DifferenceSet(a []int, b []int) []int {
	var c []int
	temp := map[int]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			c = append(c, val)
		}
	}

	return c
}
