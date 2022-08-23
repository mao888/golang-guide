package main

import (
	"fmt"
)

func main() {
	var object interface{}
	object = 1
	switch object {
	case 1:
		fmt.Printf("1") // 1
	case 2:
		fmt.Printf("2")
	case 3:
		fmt.Printf("3")
		break
	case 4:
		fmt.Printf("4")
		break
	default:
		fmt.Printf("5 ")
	}
}
