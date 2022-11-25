/**
    @author:Hasee
    @data:2022/11/25
    @note:
**/
package main

import "fmt"

func main() {
	leyangjun1 := []int{1, 3, 5, 6}
	leyangjun2 := []int{1, 3, 5}

	retDiff := DifferenceSet(leyangjun1, leyangjun2)
	fmt.Println(retDiff)
}

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
