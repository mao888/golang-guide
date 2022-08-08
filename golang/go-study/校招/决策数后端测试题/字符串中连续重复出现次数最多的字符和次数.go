/**
    @author:huchao
    @data:2022/3/11
    @note: 输出字符串中连续重复出现次数最多的字符和次数
**/
package main

import (
	"fmt"
	"math"
)

func main() {
	s := "2334"
	fmt.Println(longestRun(s))
}

//输出字符串中连续重复出现次数最多的字符和次数
func longestRun(s string) (string, int) {
	var maxVal string
	max := math.MinInt64
	maxVal = string(s[0])

	// 使用KMP算法进行匹配
	for i := 0; i < len(s); i++ {
		a := string(s[i])
		j := i + 1
		for j < len(s) {
			b := string(s[j])
			//进行匹配
			if a != b {
				//跳出当前循环
				break
			}
			j++
		}

		if j-i >= max {
			maxVal = string(s[j-1])
		}
		max = maxFunc(max, j-i)
	}
	return string(maxVal),max

}

// 判断较大的函数
func maxFunc(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}