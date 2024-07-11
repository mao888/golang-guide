package main

import (
	"fmt"
	"sync"
)

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 用 5 个携程，从 1 到 100个数，对现有 100 个数加 1 处理
func main() {
	// 定义一个长度为 100 的切片
	numbers := make([]int, 100)
	for i := range numbers {
		numbers[i] = i + 1
	}

	// 定义一个 channel，用于goroutine之间的通信
	done := make(chan bool, 5)

	var wg sync.WaitGroup
	// 定义每个携程要处理的数据
	chunkSize := len(numbers) / 5

	for i := 0; i < 5; i++ {
		// 为每个goroutine增加 wg 的计数
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			for j := start; j < start+chunkSize; j++ {
				numbers[j]++
			}
			done <- true // 发送完成信号到 channel
		}(i * chunkSize)
	}
	// 等待所有 goroutine 完成
	go func() {}()
	wg.Wait()

	// 打印结果
	for _, number := range numbers {
		fmt.Printf("%d ", number)
	}
}
