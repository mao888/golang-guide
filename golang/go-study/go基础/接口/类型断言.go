/**
    @author: HuChao
    @since: 2022/8/4
    @desc: //TODO 类型断言
**/
package main

import "fmt"

type Mover2 interface {
	move2()
}

type Pig struct {
	Name string
}

func (p Pig) move2() {
	//TODO implement me
	panic("implement me")
}

func main() {

	var n Mover2 = &Pig{Name: "pig"}
	v, ok := n.(*Pig)
	if ok {
		fmt.Println("类型断言成功")
		v.Name = "pig1" // 变量v是*Pig类型
		fmt.Println(v.Name)
	} else {
		fmt.Println("类型断言失败")
	}
}
