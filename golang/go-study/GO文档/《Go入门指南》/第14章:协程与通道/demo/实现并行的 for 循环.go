package main

import (
	"fmt"
	"sync"
)

func process(item int) {
	// 这里可以放你的处理代码
	fmt.Println(item)
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var wg sync.WaitGroup // 使用sync.WaitGroup来等待所有goroutines完成
	wg.Add(len(data))

	for _, item := range data {
		go func(item int) {
			defer wg.Done() // 当goroutine完成时，调用Done
			process(item)
		}(item)
	}

	wg.Wait() // 等待所有goroutines完成
}

//在这个示例中：
//
//1、我们使用了sync.WaitGroup来确保主函数等待所有goroutines完成。
//2、对于数据切片中的每个元素，我们都启动了一个新的goroutine来处理该元素。
//3、在goroutine内部，我们调用process()函数来处理该元素。
//4、当goroutine完成其工作后，我们调用wg.Done()来告诉WaitGroup一个goroutine已完成。
//5、在所有goroutines都启动后，我们调用wg.Wait()来等待它们都完成。
//6、这样，我们就实现了一个并行的for循环，其中每个元素都在其自己的goroutine中并行处理。
