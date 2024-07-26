package main

import "fmt"

func modifySlice(s []int) {
	s[0] = 99
	s = append(s, 4)
	fmt.Println("Inside modifySlice:", s)
}

func main() {
	slice := make([]int, 0, 6)
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	modifySlice(slice)
	fmt.Println("After modifySlice:", slice)
}
