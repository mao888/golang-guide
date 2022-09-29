/**
    @author:Hasee
    @data:2022/9/24
    @note:
**/
package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a)
	fmt.Scan(&b)
	num1 := make([]int, 0)
	num2 := make([]int, 0)
	for i := len(a) - 1; i >= 0; i-- {
		num1 = append(num1, int(a[i]-'0'))
	}
	for i := len(b) - 1; i >= 0; i-- {
		num2 = append(num2, int(b[i]-'0'))
	}
	c := add(num1, num2)
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Print(c[i])
	}

}

func add(num1, num2 []int) []int {
	t := 0
	c := make([]int, 0)
	for i := 0; i < len(num1) || i < len(num2); i++ {
		if i < len(num1) {
			t += num1[i]
		}
		if i < len(num2) {
			t += num2[i]
		}
		c = append(c, t%10)
		t = t / 10
	}
	if t != 0 {
		c = append(c, 1)
	}
	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
