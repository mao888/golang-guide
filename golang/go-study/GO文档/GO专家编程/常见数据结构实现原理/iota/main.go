/**
    @author:Hasee
    @data:2022/3/12
    @note: iota
**/
package main

import "fmt"

//iota常用于const表达式中，我们还知道其值是从零开始，const声明块中每增加一行iota值自增1

//	iota初始值为0，也即LOG_EMERG值为0，下面每个常量递增1。
type Priority int
const (
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

func main()  {
	fmt.Println(1 << 3)
	fmt.Println(1<<2)
}