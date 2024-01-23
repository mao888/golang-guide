package main

import (
	"fmt"
	gutil "github.com/mao888/mao-gutils/os"
)

func main() {
	// 测试案例： 将template中 $ 后的字符串自定义替换
	template := "chartid=$user_id & prop=$prop"

	// 声明map
	var mapStr map[string]string
	//使用make函数创建一个非nil的map，nil map不能赋值
	mapStr = make(map[string]string)
	//给已声明的map赋值
	mapStr["user_id"] = "1,2,3"
	mapStr["prop"] = "huchao"

	//res := os.Expand(template, func(s string) string {
	//	return mapStr[s]
	//})

	res := gutil.ExpandByMap(template, mapStr)
	res2 := gutil.GetComposedTemplateListExpandByMap(template, true, mapStr)
	fmt.Println(res)
	fmt.Println(res2)
}
