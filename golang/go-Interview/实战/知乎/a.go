package main

import "fmt"

//2
//3 4
//6 5 7
//4 1 8 3

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minTotal(triangle [][]int) int {
	l := len(triangle)
	if l < 1 {
		return 0
	}
	if l == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, len(triangle))
	for i, arr := range triangle {
		dp[i] = make([]int, len(arr))
	}
	result := 1<<31 - 1
	dp[0][0] = triangle[0][0] //dp[0][0]	0 0 位置所在第元素值
	dp[1][0] = triangle[1][0] + triangle[0][0]
	dp[1][1] = triangle[1][1] + triangle[0][0]

	for i := 2; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				// 包含第i行j列元素第最小路径和
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == (len(triangle[i]) - 1) {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}
	// 找到最后一行元素中，路径和最小的一个
	for _, v := range dp[len(dp)-1] {
		result = min(result, v)
	}
	return result
}

func main() {
	//triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
	var arr [][]int = [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	a := minTotal(arr)
	fmt.Println(a)
}
