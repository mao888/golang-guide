/**
    @author:Hasee
    @data:2022/3/13
    @note:
**/
package main

import "fmt"

func main()  {
	arr := []int{3,4,9,1,3,9,5}
	fmt.Println(findKDistantIndices(arr,9,1))
}

func findKDistantIndices(nums []int, key int, k int) []int {
	var res []int
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := max(0, i-k); j < min(n, i+k+1); j++ {
			if nums[j] == key {
				res = append(res, i)
				break
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}