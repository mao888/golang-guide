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

// 三个 goroutine，两个goroutine计算（不打印），一个加到 1 万，一个加到 2 万
// 第三个等他们两个算完，将他们两个的和打印出来
func main() {
	// 三个 goroutine，两个goroutine计算（不打印），一个加到 1 万，一个加到 2 万
	// 第三个等他们两个算完，将他们两个的和打印出来
	var wg sync.WaitGroup
	sum1Chan := make(chan int)
	sum2Chan := make(chan int)

	// 第一个
	wg.Add(1)
	go func() {
		defer wg.Done()
		sum := 0
		for i := 1; i <= 10000; i++ {
			sum += i
		}
		sum1Chan <- sum
	}()

	// 第二个
	wg.Add(1)
	go func() {
		defer wg.Done()
		sum := 0
		for i := 1; i <= 20000; i++ {
			sum += i
		}
		sum2Chan <- sum
	}()

	// 第三个
	wg.Add(1)
	go func() {
		defer wg.Done()
		sum1 := <-sum1Chan
		sum2 := <-sum2Chan
		fmt.Println("sum: ", sum1+sum2)
	}()
	wg.Wait()
}
