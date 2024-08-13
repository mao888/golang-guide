package main

import (
	"fmt"
	"sync"
)

// Counter 一个线程安全的计数器结构体
type Counter struct {
	mu    sync.Mutex
	count int
}

// Increment 增加计数
func (c *Counter) Increment() {
	c.mu.Lock() // 加锁，确保线程安全
	c.count++
	c.mu.Unlock() // 解锁
}

// Value 获取当前计数值
func (c *Counter) Value() int {
	c.mu.Lock()         // 加锁，确保读取时的线程安全
	defer c.mu.Unlock() // defer确保方法结束时解锁
	return c.count
}

func main() {
	counter := &Counter{}

	var wg sync.WaitGroup

	// 启动多个 goroutine 并发地增加计数器
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Value:", counter.Value())
}
