/**
    @author:Hasee
    @data:2022/9/24
    @note:
**/

//https://blog.csdn.net/a568283992/article/details/119698329

package main

import (
	"fmt"
)

func main() {
	a := big_integer.ValueOf(6782613786431)
	b := big_integer.ValueOf(-678261378231)
	res := a.Add(b)
	fmt.Println(res.String())
}
