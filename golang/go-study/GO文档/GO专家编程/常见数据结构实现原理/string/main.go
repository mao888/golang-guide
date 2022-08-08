/**
    @author:Hasee
    @data:2022/3/12
    @note: string
**/
package main

import "fmt"

func main()  {
	var a []string
	a = append(a, "sdf","sdff","sd")
	fmt.Println(concatstrings(a))
	//str := "Str1" + "Str2" + "Str3"
	fmt.Println(concatstrings(a))
}

func concatstrings(a []string) string { // 字符串拼接
	length := 0        // 拼接后总的字符串长度

	for _, str := range a {
		length += len(str)
	}

	s, b := rawstring(length) // 生成指定大小的字符串，返回一个string和切片，二者共享内存空间

	for _, str := range a {
		copy(b, str)    // string无法修改，只能通过切片修改
		b = b[len(str):]
	}

	return s
}
// 因为string是无法直接修改的，
// 所以这里使用rawstring()方法初始化一个指定大小的string，
// 同时返回一个切片，二者共享同一块内存空间，
// 后面向切片中拷贝数据，也就间接修改了string。
