package main

import (
	"fmt"
	"strings"
)

func main() {
	dateString := "2023-11-09"

	// 使用空格分割字符串
	parts := strings.Split(dateString, " ")

	// 提取日期部分
	if len(parts) >= 1 {
		day := parts[0]
		fmt.Println("截取到天:", day)
	} else {
		fmt.Println("日期字符串格式不正确")
	}
}
