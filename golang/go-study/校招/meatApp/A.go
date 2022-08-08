/**
    @author: huchao
    @since: 2022/8/3
    @desc: //TODO
**/
package main

import "fmt"

func main() {
	a := decha(5)
	fmt.Println(a)

}

func anser(x []int) []int {
	//a := decha(5)
	c := sort(x)
	return c
}

func decha(input int) []int {
	a := make([]int, 10)
	a[9] = 65

	for i := 8; i >= 0; i-- {
		a[i] = a[i+1] + input
	}

	return a
}

func sort(x []int) []int {
	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i] > x[j] {
				temp := x[i]
				x[i] = x[j]
				x[j] = temp
			}
		}
	}
	return x
}
