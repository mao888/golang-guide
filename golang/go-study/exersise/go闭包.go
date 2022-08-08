/**
    @author:huchao
    @data:2022/2/11
    @note:
**/
package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int { base += i
		return base }
	sub := func(i int) int { base -= i
		return base }
	return add, sub
}
func main() {
	f1,f2 := calc(10)
	fmt.Println(f1(1),f2(2))	// 11  9
	fmt.Println(f1(3),f2(4))	// 12  8
	fmt.Println(f1(5),f2(6))	// 13  7
	fmt.Println(f1(7),f2(8))	// 14  6
}
