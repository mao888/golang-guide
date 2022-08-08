/**
    @author:Hasee
    @data:2022/3/19
    @note:
**/
package main

import "fmt"

func main() {
	s := "Escape"
	fmt.Println(s)	// 很多函数参数为interface类型，比如fmt.Println(a …interface{})，编译期间很难确定其参数的具体类型，也会产生逃逸。

}