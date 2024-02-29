package main

import (
	"fmt"
	"sync"
)

var (
	concurrentNum2 = 3
)

type Semaphore struct {
	ch       chan struct{}
	wg       sync.WaitGroup
	maxConns int
}

func main() {
	repositoryClean2()
}

func repositoryClean2() {
	sem := NewSemaphore2(concurrentNum2)
	for i := 0; i < 10; i++ {
		sem.Acquire()
		fmt.Printf("goroutine %d is running\n", i)
		go func(i int) {
			defer sem.Release()
			defer fmt.Printf("goroutine %d completed\n", i)
			doGcCommand2()
		}(i)
	}
	sem.Wait()
}

func NewSemaphore2(maxConns int) *Semaphore {
	return &Semaphore{
		ch:       make(chan struct{}, maxConns),
		maxConns: maxConns,
	}
}

func (s *Semaphore) Acquire() {
	s.wg.Add(1)

	s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.ch
	s.wg.Done()
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}

func doGcCommand2() {
	fmt.Println("   doGcCommand executed")
}
