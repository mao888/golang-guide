package main

import (
	"fmt"
	"strconv"
)

// 浮点数 1.0 输出后保存小数位
func main() {
	var f float64 = 1.0
	fmt.Println(f) // 1
	value := strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Println(value) // 1.00
}
