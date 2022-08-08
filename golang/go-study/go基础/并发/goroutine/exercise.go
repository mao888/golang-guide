package main

import (
	"fmt"
	"sync"
)

var wg3 sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg3.Add(1)
		go func(i int) {
			defer wg3.Done()
			fmt.Println(i)
		}(i)
	}
	wg3.Wait()
}
