/**
    @author: huchao
    @since: 2022/8/3
    @desc: //TODO 类型与接口的关系
**/
package main

import "fmt"

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

//	海尔
type haier struct {
	dryer
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

var h haier
var wa WashingMachine

func main() {

	wa = h
	wa.wash()
	wa.dry()

	var haier = haier{}
	var s = WashingMachine(haier)
	s.wash()
	s.dry()
	haier.dry()
	haier.dryer.dry()
	haier.wash()
}
