/**
    @author: huchao
    @since: 2022/8/2
    @desc: //TODO 计算int64随机数各位数和
**/
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

//	开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
func createRandNumber() <-chan int64 {
	jobChain := make(chan int64, 10)
	go func() {
		for i := 0; i < 100; i++ {
			jobChain <- rand.Int63()
		}
		close(jobChain)
	}()
	return jobChain
}

//	开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
func receiveJobChain(jobChain <-chan int64, resultChain chan<- int) {
	defer wg.Done()
	for v := range jobChain {
		sum := 0
		for v > 0 {
			sum += int(sum % 10)
			sum /= 10
		}
		resultChain <- sum
	}
}

func main() {
	jobChain := createRandNumber()

	resultChain := make(chan int, 100)
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go receiveJobChain(jobChain, resultChain)
	}

	go func() {
		wg.Wait()
		close(resultChain)
	}()
	
	for v := range resultChain {
		fmt.Println(v)
	}
}
