package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

// 文档：https://www.liwenzhou.com/posts/Go/ratelimit/#autoid-0-2-0
// github：https://github.com/uber-go/ratelimit

func main() {
	rl := ratelimit.New(100) // per second

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}

	// Output:
	// 0 0
	// 1 10ms
	// 2 10ms
	// 3 10ms
	// 4 10ms
	// 5 10ms
	// 6 10ms
	// 7 10ms
	// 8 10ms
	// 9 10ms
}
