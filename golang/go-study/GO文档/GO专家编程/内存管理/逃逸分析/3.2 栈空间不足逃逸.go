/**
    @author:Hasee
    @data:2022/3/19
    @note:
**/
package main

//我们发现当切片长度扩大到10000时就会逃逸。
//
//实际上当栈空间不足以存放当前对象时或无法判断当前切片长度时会将对象分配到堆中。

func Slice() {
	s := make([]int, 1000, 1000)

	for index, _ := range s {
		s[index] = index
	}
}

func main() {
	Slice()
}