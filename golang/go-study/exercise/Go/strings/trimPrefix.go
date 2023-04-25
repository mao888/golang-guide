package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "10110:[100186,100187]:[a,w,s,d]"
	str = strings.TrimPrefix(str, "[")
	str = strings.TrimSuffix(str, "]")
	fmt.Println(str)
}
