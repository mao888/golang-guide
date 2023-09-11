package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := " 测试第一条ad s "
	charCount := utf8.RuneCountInString(str)            // 用于获取字符串的字符数，而不是字节数
	fmt.Println("len str", len(str))                    // 21 = 1 + 15 + 2 + 1 + 1 + 1
	fmt.Printf("字符串 \"%s\" 的字符数为 %d\n", str, charCount) // 11
	trimSpaceStr := strings.TrimSpace(str)              // 去除字符串首尾的空白字符

	fmt.Println("len trimSpaceStr", len(trimSpaceStr))                                      // 15 + 4 = 19
	fmt.Printf("字符串 \"%s\" 的字符数为 %d\n", trimSpaceStr, utf8.RuneCountInString(trimSpaceStr)) // 9
}
