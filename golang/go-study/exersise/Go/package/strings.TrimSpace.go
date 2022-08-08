package main

import (
	"fmt"
	"strings"
)

/**
[]byte中每一个byte的默认值是0，而空格的值是32；按照实验的结果，TrimSpace可以剪掉byte为32的（空格），但不能剪掉byte为0的
返回将s前后端所有空白（unicode.IsSpace指定）都去掉的字符串。
*/

func testa() {
	buf := make([]byte, 10)
	buf[0] = ' '
	buf[1] = ' '
	buf[2] = 'b'
	buf[3] = 'b'
	buf[4] = ' '
	//buf[5] = ' '
	//buf[6] = 'b'
	// buf[7] = ' '
	// buf[8] = ' '
	buf[9] = ' '
	fmt.Println(buf)
	fmt.Printf("%s\n", strings.TrimSpace(string(buf)))
	fmt.Printf("%d\n", len(strings.TrimSpace(string(buf))))
}

func main() {
	testa()

	//testb()
}

func testb() {
	s := "Hello world hello world"
	ret := strings.TrimSpace(s)
	fmt.Println(ret)
	fmt.Println(len(ret))
}
