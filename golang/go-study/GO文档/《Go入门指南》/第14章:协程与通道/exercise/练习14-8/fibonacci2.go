package main

import "fmt"

func main() {
	pos := 4
	result, pos := fibonacci2(pos)
	fmt.Printf("the %d-th fibonacci number is: %d\n", pos, result)
	pos = 10
	result, pos = fibonacci2(pos)
	fmt.Printf("the %d-th fibonacci number is: %d\n", pos, result)
}

func fibonacci2(n int) (val, pos int) {
	if n <= 1 {
		val = 1
	} else {
		v1, _ := fibonacci2(n - 1)
		v2, _ := fibonacci2(n - 2)
		val = v1 + v2
	}
	pos = n
	return
}
