package main

import "fmt"

func main() {
	list := []string{"a", "b", "c", "d", "e"}
	slice := list[2:4:6]
	fmt.Println(slice)
}
