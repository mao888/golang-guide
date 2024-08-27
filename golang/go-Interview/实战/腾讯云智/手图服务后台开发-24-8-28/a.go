package main

import (
	"fmt"
	"sync"
)

// golang实现3个goroutine交替打印1~100。例如：
// Goroutine 1: 1
// Goroutine 2: 2
// Goroutine 3: 3
// Goroutine 1: 4
// Goroutine 2: 5
// Goroutine 3: 6
// Goroutine 1: 7
// Goroutine 2: 8
// Goroutine 3: 9
// Goroutine 1: 10
// Goroutine 2: 11
// Goroutine 3: 12
func main() {
	var wg sync.WaitGroup
	printChan := make(chan int, 1)

	goroutines := 3
	totalNumbers := 100

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for {
				num, ok := <-printChan
				if !ok {
					return
				}
				if num > totalNumbers {
					close(printChan)
					return
				}
				if (num-1)%goroutines == id {
					fmt.Printf("Goroutine %d: %d\n", id+1, num)
					printChan <- num + 1
				} else {
					printChan <- num
				}
			}
		}(i)
	}

	printChan <- 1
	wg.Wait()
}
