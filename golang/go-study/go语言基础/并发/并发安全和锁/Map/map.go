/**
    @author: huchao
    @since: 2022/8/2
    @desc: //TODO sync.Map
**/
package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
	Go 语言中内置的 map 不是并发安全的,如下代码
*/

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			set(key, i)
			fmt.Printf("k=:%v,v=:%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
