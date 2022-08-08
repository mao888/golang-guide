/**
    @author:Hasee
    @data:2022/4/4
    @note:
**/
package main

import (
	"fmt"
	"time"
)

func main() {
	var a string
	var b string
	var c string
	a = "2021-04-30 03:21:18"
	b = "2022-04-30 03:21:18"
	c = "2022-05-01 03:21:18"
	baseFormat := "2006-01-02 15:04:05"      //常规类型
	a1, _ := time.Parse(baseFormat, a)       // a1: "2021-04-30 03:21:18"
	b1, _ := time.Parse(baseFormat, b)       // b1: "2022-04-30 03:21:18"
	c1, _ := time.Parse(baseFormat, c)       // c1: "2022-05-01 03:21:18"
	fmt.Println(a1)                          // 2021-04-30 03:21:18 +0000 UTC
	fmt.Println(b1)                          // 2022-04-30 03:21:18 +0000 UTC
	fmt.Println(c1)                          // 2022-05-01 03:21:18 +0000 UTC
	fmt.Println(b1.Before(c1))               // b1 在 c1前 true
	fmt.Println(b1.Format(baseFormat))       // 2022-04-30 03:21:18
	fmt.Println(a1.YearDay())                // 120
	fmt.Println(b1.YearDay())                // 120
	fmt.Println(c1.YearDay())                // 121
	fmt.Println(c1.YearDay() - b1.YearDay()) // 1
	fmt.Println(a1.UnixNano())               // 1619752878000000000
}
