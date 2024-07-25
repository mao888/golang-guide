package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//	type ListNode struct {
//		Val  int
//		Next *ListNode
//	}
//
// m 个生产者，n 个消费者，用 channel 进行数据通信，消费函数为 func work（msg int） error，当消费捕获到 error 的时候，关闭所有的生产者和消费者。
// 在Go语言中，使用channel进行生产者和消费者之间的数据通信是一种常见的模式。
// 如果你希望在一个消费者捕获到错误时关闭所有的生产者和消费者，你可以采用一些全局控制机制，
// 比如一个共享的错误通道（error channel）、一个关闭信号（如context）或者一个共享的标志位（使用sync.WaitGroup和sync.Mutex来安全地访问和修改）。

// 模拟工作函数，可能返回错误
func work(ctx context.Context, msg int) error {
	// 模拟处理时间
	time.Sleep(time.Millisecond * 100)
	// 假设当消息为负数时，产生错误
	if msg < 0 {
		return errors.New("negative value")
	}
	fmt.Printf("Processed %d\n", msg)
	return nil
}

// 生产者
func producer(ctx context.Context, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- 1                            // 这里为了简单，我们始终发送1
			time.Sleep(time.Millisecond * 200) // 模拟生产间隔
		}
	}
}

// 消费者
func consumer(ctx context.Context, wg *sync.WaitGroup, ch <-chan int, done chan<- struct{}) {
	defer wg.Done()
	for msg := range ch {
		if err := work(ctx, msg); err != nil {
			fmt.Println("Error:", err)
			// 发送关闭信号
			close(done)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	ch := make(chan int, 10)
	done := make(chan struct{})

	// 启动生产者
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go producer(ctx, &wg, ch)
	}

	// 启动消费者
	wg.Add(1)
	go consumer(ctx, &wg, ch, done)

	// 等待关闭信号或所有goroutine完成
	go func() {
		wg.Wait()
		close(ch) // 关闭channel以通知生产者结束
	}()

	// 等待消费者发送的关闭信号
	<-done
	cancel() // 取消所有子上下文

	// 注意：这里不再等待wg.Wait()，因为我们已经通过done通道得知发生了错误
	fmt.Println("Shutting down...")
}

//Context: 使用context来优雅地关闭生产者和消费者。当消费者捕获到错误时，可以通过关闭done通道来触发关闭流程。
//WaitGroup: 用于等待所有的goroutine完成。然而，在本例中，当done通道被关闭时，我们实际上不再需要等待wg.Wait()，因为错误已经发生，系统需要立即开始关闭流程。
//Channel关闭: 在消费者中不直接关闭生产者向其中发送数据的channel（ch），而是由主goroutine在wg.Wait()后关闭。这是因为如果消费者关闭了这个channel，生产者可能会尝试向已关闭的channel发送数据，导致panic。
//错误处理: 消费者中一旦捕获到错误，就关闭done通道并退出。主goroutine监测到这个信号后，通过cancel()取消所有子上下文，并开始清理流程。
