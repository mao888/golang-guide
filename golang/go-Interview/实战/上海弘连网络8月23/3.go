package main

import (
	"fmt"
)

var item = "hello"

func main() {
	v := item
	v[0] = 'a'
	fmt.Printf("%s", item)
}
