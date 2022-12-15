package main

// 模拟并行计算  用协程并行输出
// 问题：并发求和？
func main() {

}

func SubSum(start int, len int) {
	var ssum int64
	ssum = 0
	for i := 1; i <= len; i++ {
		ssum += int64(start)
		start++
	}

}
