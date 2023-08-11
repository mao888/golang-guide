package main

import (
	"fmt"
	"sync"
)

func searchInArray(arr []int, target int, startIndex int, endIndex int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	for i := startIndex; i < endIndex; i++ {
		if arr[i] == target {
			resultChan <- i
			return
		}
	}

	resultChan <- -1
}

func main() {
	arr := make([]int, 1000000) // 假设有一个包含1000000个元素的数组
	target := 42                // 要查找的目标元素

	numWorkers := 4 // 协程数量
	wg := sync.WaitGroup{}
	resultChan := make(chan int, numWorkers)

	arrayChunkSize := len(arr) / numWorkers

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		startIndex := i * arrayChunkSize
		endIndex := (i + 1) * arrayChunkSize
		if i == numWorkers-1 {
			endIndex = len(arr)
		}
		go searchInArray(arr, target, startIndex, endIndex, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	foundIndex := -1
	for idx := range resultChan {
		if idx != -1 {
			foundIndex = idx
			break
		}
	}

	if foundIndex != -1 {
		fmt.Printf("元素 %d 在数组中的索引为 %d\n", target, foundIndex)
	} else {
		fmt.Printf("未找到元素 %d\n", target)
	}
}
