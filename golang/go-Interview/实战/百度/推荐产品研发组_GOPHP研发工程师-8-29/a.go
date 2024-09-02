package main

import (
	"fmt"
	"math"
)

//一个n米深的井，每分钟只能爬u米，每次爬之前都要休息一分钟，休息期间会下滑d米，求爬出井需要多少分钟？
//输入n，u，d，输出所需分钟

func minute(n, u, d int) int {
	// 爬升净值
	netClimb := u - d
	if n <= u {
		return 1
	}
	// 除最后一次爬升的循环次数
	cycles := n/netClimb - 1
	// 总时长：每次循环 2 分钟，加上最后一次爬升1分钟
	total := cycles*2 + 1
	return total
}

func calculateMinutes(n, u, d int) int {
	// 每次实际上升的高度
	actualClimb := u - d
	// 除最后一次外，需要爬升的次数
	climbsNeeded := math.Ceil(float64(n-u) / float64(actualClimb))
	// 总分钟数：爬升次数*2（每次爬升和休息）+ 最后一次爬升
	totalMinutes := int(climbsNeeded*2) + 1
	return totalMinutes
}

func main() {
	n := 10
	u := 2
	d := 1
	fmt.Println(minute(n, u, d))
	fmt.Println(calculateMinutes(n, u, d))

	fmt.Println("--------------------------------")
	n = 10
	u = 3
	d = 1
	fmt.Println(minute(n, u, d))
	fmt.Println(calculateMinutes(n, u, d))
}
