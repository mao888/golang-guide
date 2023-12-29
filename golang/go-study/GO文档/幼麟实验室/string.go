package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

// 视频: https://www.bilibili.com/video/BV1ff4y1m72A/?spm_id_from=333.999.0.0&vd_source=8321160752e4f07c473e11ebc0dd0a28

var s2 string = " eggo世界"

const s3 = " eggo世界"

func main() {
	var s1 string = " eggo世界"
	// 字符串的长度
	fmt.Println(len(s1)) // 11
	println(utf8.RuneCountInString(s1))
	fmt.Println(s1[2])         // 103
	fmt.Printf("%c\n", s1[2])  // g
	fmt.Printf("%s\n", s1[2])  // %!s(uint8=103)
	fmt.Printf("%s\n", s1[:2]) // e

	// 修改字符串 []byte 会开辟新的内存空间
	//s1[2] = 'a' // cannot assign to s1[2]
	fmt.Println(s1)

	bs := []byte(s1)
	fmt.Println(bs) // [32 101 103 103 111  228 184 150  231 149 140]
	fmt.Println(bs) // [32 101 97 103 111 228 184 150 231 149 140]
	fmt.Printf("%c\n", bs[2])

	// 修改字符串 []rune 会开辟新的内存空间
	runes := []rune(s1)
	fmt.Println(runes) // [32 101 103 103 111 19990 30028]
	runes[5] = 'a'
	fmt.Printf("%c\n", runes[2])
	fmt.Println(string(runes)) //  eggoa界
	runes[5] = 32
	fmt.Println(string(runes)) //  eggo 界

	// unsafe 包 让slice依然使用原来字符串指向的这段内存
	// 将字符串转换为[]byte，获取底层字节数组的指针
	// 将字符串的头部信息转换为 reflect.StringHeader 结构
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s1))

	// 构造一个 reflect.SliceHeader 结构，其中 Data 指向字符串的底层字节数组，
	// Len 表示切片的长度，Cap 表示切片的容量，这里我们将 Len 和 Cap 设置为字符串的长度
	// 将 reflect.SliceHeader 转换为 []byte 切片
	byteSlice := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: strHeader.Data,
		Len:  strHeader.Len,
		Cap:  strHeader.Len,
	}))

	// 打印底层字节数组
	fmt.Printf("底层字节数组: %v\n", byteSlice)

	strHeader.Len = 5
	fmt.Println(len(s1))

	// 底层地址
	fmt.Printf("%p\n", &s1)                                                     // 0x14000110210
	fmt.Printf("%p\n", strHeader)                                               // 0x14000110210
	fmt.Printf("%p\n", &byteSlice)                                              // 0x1400011e060
	fmt.Printf("%x\n", strHeader.Data)                                          // 10258a13d
	fmt.Printf("%x\n", (*reflect.SliceHeader)(unsafe.Pointer(&byteSlice)).Data) // 104c3e17d
}

func a() {
	const name int = iota

	a := 1
	println(a)
}
