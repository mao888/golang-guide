package main

import (
	"fmt"
	gutil "github.com/mao888/go-utils/interfaces"
)

func main() {
	var name interface{} = "aaa"
	fmt.Println(gutil.JudgeType(name))
	fmt.Println(gutil.JudgeTypeByReflect(name))
}
