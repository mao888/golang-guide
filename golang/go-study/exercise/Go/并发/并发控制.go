package main

import (
	"fmt"
	"sync"
)

var (
	// 指定允许的最大并发数量
	concurrentNum = 1
)

type limitSyncGroup struct {
	c  chan struct{} // 一个带缓冲的 channel，其缓冲大小限制了并发数量
	wg *sync.WaitGroup
}

func main() {
	repositoryClean()
}
func repositoryClean() {
	// 创建了一个限制了并发数量的信号量
	var wait = NewSemaphore(concurrentNum)
	wait.Add(10)
	for i := 0; i < 10; i++ {

		//wait.Add(1)
		fmt.Printf("goroutine %d is running\n", i)
		go func(i int) {
			defer fmt.Printf("goroutine %d completed\n", i)
			defer wait.Done()
			doGcCommand()
		}(i)
	}
	wait.Wait()
}

func NewSemaphore(maxSize int) *limitSyncGroup {
	return &limitSyncGroup{
		c:  make(chan struct{}, maxSize),
		wg: new(sync.WaitGroup),
	}
}
func doGcCommand() {
	fmt.Println("   doGcCommand 执行了")
}

func (s *limitSyncGroup) Add(delta int) {
	s.wg.Add(delta)
	s.c <- struct{}{}
}
func (s *limitSyncGroup) Done() {
	<-s.c
	s.wg.Done()
}
func (s *limitSyncGroup) Wait() {
	s.wg.Wait()
}
