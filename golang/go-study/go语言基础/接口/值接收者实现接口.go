/**
    @author: huchao
    @since: 2022/8/3
    @desc: //TODO 值接收者实现接口
**/
package main

import "fmt"

// Mover 定义一个接口类型
type Mover interface {
	Move()
}

// Dog 狗结构体类型
type Dog struct{}

// Move 使用值接收者定义Move方法实现Mover接口
func (d Dog) Move() {
	fmt.Println("狗会走🐶")
}

var x Mover // 声明一个Mover类型的变量x

var d1 = Dog{} // d1是Dog类型

var d2 = &Dog{} // d2是Dog指针类型

func main() {
	x = d1 // 可以将d1赋值给变量x
	x.Move()

	x = d2 // 也可以将d2赋值给变量x
	x.Move()
}

/*
从上面的代码中我们可以发现，
使用值接收者实现接口之后，
不管是结构体类型还是对应的结构体指针类型的变量都可以赋值给该接口变量
*/
