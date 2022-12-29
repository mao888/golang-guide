// loopmutex project selectsort.go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//1.这个程序总的功能是计算1到100000（LEN×STEP）的和，计算被分为STEP步（段）进行，通过调用函数subsum()计算每段之和，然后相加
//
//2.函数subsum()计算从start开始长度为len的数列之和
//
//3.语句"fmt.Printf("%d\n", (int64(LEN*STEP) * (int64(LEN*STEP) + 1) / int64(2)))"打印输出一个正确的结果作为参考值
//
//4.为了知道程序当前有几个goroutine在运行，需要使用包"runtime"，其中的方法runtime.NumGoroutine()返回正在运行的goroutine的数量，需要注意的是main()本身也是一个goroutine
//
//5.使用包"time"中的方法time.Sleep()，让自身的goroutine休眠，代入参数指定休眠0.1秒，因为需要等待所有其他goroutine都执行完之后程序才能结束
//
//6.使用互斥锁mu来锁住变量sum（参见程序），需要使用包"sync"

const START = 1
const LEN = 100
const STEP = 1000

var (
	mu  sync.Mutex
	sum int64
)

func main() {
	fmt.Printf("%d\n", (int64(LEN*STEP) * (int64(LEN*STEP) + 1) / int64(2)))

	sum = 0
	start := START
	for i := 1; i <= STEP; i++ {
		go subsum(start, LEN)
		start += LEN
	}

	// runtime.NumGoroutine()返回正在运行的goroutine的数量
	for runtime.NumGoroutine() > 1 {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("%d\n", sum)
}

func subsum(start int, len int) {
	var ssum int64
	ssum = 0
	for i := 1; i <= len; i++ {
		ssum += int64(start)
		start++
	}

	mu.Lock()
	sum += ssum
	mu.Unlock()
}
