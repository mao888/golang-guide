/**
    @author: huchao
    @since: 2022/8/3
    @desc: //TODO 指针接收者实现接口
**/
package main

import "fmt"

// Cat 猫结构体类型
type Cat struct{}

// Move 使用指针接收者定义Move方法实现Mover接口
func (c *Cat) Move() {
	fmt.Println("猫会动")
}

func main() {
	var c1 = &Cat{} // c1是*Cat类型

	x = c1 // 可以将c1当成Mover类型
	x.Move()

	/*
		// 下面的代码无法通过编译
		var c2 = Cat{} // c2是Cat类型
		x = c2         // 不能将c2当成Mover类型
	*/
}

/*
由于Go语言中有对指针求值的语法糖，
对于值接收者实现的接口，无论使用值类型还是指针类型都没有问题。
但是我们并不总是能对一个值求址，所以对于指针接收者实现的接口要额外注意。
*/