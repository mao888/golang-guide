package main

import (
	"fmt"
)

func main() {
	//初始化map[string][]string
	funcMap := map[string][]string{
		"1111":  []string{"2222", "333333"},
		"11115": []string{"2ooo222", "33ooo3333"},
	}
	fmt.Println(funcMap) //替换之前的打印   map[1111:[2222 333333] 11115:[2ooo222 33ooo3333]]
	funcMap1 := map[string]string{
		"2ooo222": "33333fefe", //切片中的数值按照这个映射替换
	}

	for _, funcSlice := range funcMap { //map遍历
		for loc, funcName := range funcSlice { //[]string遍历
			if v, ok := funcMap1[funcName]; ok { //替换切片中数值
				funcSlice[loc] = v
			}
		}
	}
	fmt.Println(funcMap) //替换之后的打印  map[1111:[2222 333333] 11115:[33333fefe 33ooo3333]]
}

//funcSlice 这是一个string的切片,底层和funcMap[key]的值时共用一个地址的，所以直接修改funcSlice就能修改map中切片的值。就类似于如下代码：
//
//g:=[]string{11,2,3,4,45,6}
//c=g[1:]
//g和c目前有部分是公用的，修改c就能修改g，前提是不能扩容c。
