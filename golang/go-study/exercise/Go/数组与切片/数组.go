/**
    @author:Hasee
    @data:2022/6/28
    @note:
**/
package main

import "fmt"

func main() {
	var a1 = []int{5, 1, 4}
	var a2 = []int{3}
	for _, i2 := range a2 {
		a1 = append(a1, i2)
	}
	fmt.Println(a1)
}
