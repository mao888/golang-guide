package main

import (
	"fmt"
	gutil "github.com/mao888/mao-gutils/interfaces"
)

func main() {
	var name interface{} = "aaa"
	fmt.Println(gutil.JudgeType(name))
	fmt.Println(gutil.JudgeTypeByReflect(name))
}
