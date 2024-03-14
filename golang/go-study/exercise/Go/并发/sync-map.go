package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	// 1. 写入
	m.Store("qcrao", 18)
	m.Store("stefno", 20)

	// 2. 读取
	age, _ := m.Load("qcrao")
	fmt.Println(age.(int))

	// 3. 遍历
	m.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)
		fmt.Println(name, age)
		return true
	})

	// 4. 删除
	m.Delete("qcrao")
	age, ok := m.Load("qcrao")
	fmt.Println(age, ok)

	// 5. 读取或写入
	m.LoadOrStore("stefno", 100)
	age, _ = m.Load("stefno")
	fmt.Println(age)
}
