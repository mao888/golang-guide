package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 程序运行时候启动2个线程，2个线程都会访问相同的一串数据，数据的数量和类型是不确定的，请编写一段程序，可以让2个线程在无锁的情况下，访问修改添加删除这段数据和数据里的任意数据类型。
func main() {
	var data sync.Map
	var counter int32

	data.Store("key1", "value1")
	data.Store("key2", "value2")

	// first goroutine 用于添加新数据
	go func() {
		for i := 0; i < 10; i++ {
			key := fmt.Sprintf("key%d", atomic.AddInt32(&counter, 1))
			value := fmt.Sprintf("value%d", i)
			data.Store(key, value)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// second goroutine	用于读取和打印数据
	go func() {
		for i := 0; i < 10; i++ {
			key := fmt.Sprintf("key%d", i)
			if value, ok := data.Load(key); ok {
				fmt.Printf("Key: %s, Value: %v\n", key, value)
			} else {
				fmt.Printf("Key: %s not found\n", key)
			}
			time.Sleep(150 * time.Millisecond)
		}
	}()

	// 确保 goroutine 有时间运行
	time.Sleep(2 * time.Second)

	data.Range(func(key, value interface{}) bool {
		fmt.Printf("Final Key: %v, Final Value: %v\n", key, value)
		return true
	})
}
