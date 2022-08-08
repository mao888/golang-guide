/**
    @author:Hasee
    @data:2022/7/7
    @note:
**/

package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const arraySize = 10

type CircularBuffer struct {
	data    [arraySize]int
	pointer int
}

//只实现了CircularBuffer环形缓冲队列的基本方法
func (b *CircularBuffer) InsertValue(i int) {
	if b.pointer == len(b.data) {
		b.pointer = 0
	}
	b.data[b.pointer] = i
	b.pointer += 1
}

func (b *CircularBuffer) GetValues() [arraySize]int {
	return b.data
}

func (b *CircularBuffer) GetValuesFromPosition(i int) ([arraySize]int, bool) {
	var out [arraySize]int

	if i >= len(out) {
		return out, false
	}

	for u := 0; u < len(out); u++ {
		if i >= len(b.data) {
			i = 0
		}
		out[u] = b.data[i]
		i += 1
	}
	return out, true
}

func TestCircularBuffer(t *testing.T) {
	var cb CircularBuffer
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	//生成三轮随机数，来测试这个CircularBuffer
	for i := 0; i < 3*len(cb.data); i++ {
		cb.InsertValue(random.Intn(arraySize))
		fmt.Println(cb, cb.pointer, (i+1)%arraySize)
		//判断时，要排除在pointer最大时，因为此时，余数为0，而pointer为数组长度
		if cb.pointer != arraySize && cb.pointer != (i+1)%arraySize {
			t.Fail()
		}
		//其它几个方法，待测试
	}
}


