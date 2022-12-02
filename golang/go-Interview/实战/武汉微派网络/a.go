package main

import "fmt"

// 从字符串中提取数字
func main() {
	str := "a123a"
	ss := FindNum(str)
	fmt.Println(ss)
}

func FindNum(char string) int {
	// i记录循环次数
	// k记录提取的单个数的值
	// flag记录当前是否为数字
	// sum记录a[]的下标，存放数字到a数组中
	//var a []int
	var flag int
	var m int
	var sum int
	var k int
	flag = 0
	sum = 0
	m = 0

	for i := 0; i < len(char); i++ {
		// 当前位为数字，则放入k中
		if char[i] >= '0' && char[i] <= '9' {
			flag = 1
			k = int(char[i] - '0')
		} else {
			// flag = 1 ,说明前一位是数字，而现在已经不是数字
			// 把计算的sum放入数组中，一组连在一起的数字判断完成
			if flag == 1 {
				m = m + 1
				//a[m] = sum
			}
			// 复位
			flag = 0
			sum = 0
		}
		// 若为数字，则计算连在一起的数字的值
		if flag == 1 {
			sum = sum * 10
			sum += k
		}
	}
	return sum
}
