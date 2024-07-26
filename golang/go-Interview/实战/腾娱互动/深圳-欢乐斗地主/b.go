package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			fmt.Println(i)
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}
