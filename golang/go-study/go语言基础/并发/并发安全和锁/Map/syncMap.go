/**
    @author: huchao
    @since: 2022/8/2
    @desc: //TODO 并发读写sync.Map
**/
package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 并发安全对map
var sm = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	// 对sm执行20个并发的读写操作
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			sm.Store(key, i)         // 存储key-value
			value, _ := sm.Load(key) // 根据key取值
			fmt.Printf("k:=%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
