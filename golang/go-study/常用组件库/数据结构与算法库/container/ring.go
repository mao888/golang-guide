package main

import (
	"container/ring"
	"fmt"
)

// ring实现了环形链表的操作。
// http://doc.golang.ltd/

func main() {
	ring := ring.New(5)
	fmt.Println(ring)
}
