package main

import (
	"fmt"
	o "github.com/mao888/go-utils/os"
)

func main() {
	template := "chartid=$user_id & prop=$prop"

	// 声明map
	var mapStr map[string]string
	//使用make函数创建一个非nil的map，nil map不能赋值
	mapStr = make(map[string]string)
	//给已声明的map赋值
	mapStr["user_id"] = "123"
	mapStr["prop"] = "huchao"

	//res := os.Expand(template, func(s string) string {
	//	return mapStr[s]
	//})

	res := o.ExpandByMap(template, mapStr)
	fmt.Println(res)
}
