/**
    @author:Huchao
    @data:2023/1/6
    @note:数据结构之线性表--顺序表 测试
**/
package main

import (
	"fmt"
	"testing"
)

func TestSqList(t *testing.T) {
	var li SqList
	// 初始化
	li.InitList(4)
	// 判空
	fmt.Println(li.ListEmpty()) // true
	// 判满
	fmt.Println(li.ListFul()) // false
	// 定义一个Struct类型
	type s struct {
		name string
		age  int
	}
	student1 := s{name: "abc", age: 10}
	student2 := s{name: "efg", age: 10}
	// 插入元素
	li.ListInsert(0, student1)
	li.ListInsert(1, student2)
	// 判空
	fmt.Println(li.ListEmpty()) // false
	// 插入元素
	li.ListInsert(2, 1000)
	li.ListInsert(3, "GOGO")

	// 遍历
	li.TraverseList() // {abc 10} {efg 10} 1000 GoGO
	// 获取长度
	fmt.Println(li.ListLength()) // 4
	// 判满
	fmt.Println(li.ListFul()) // true
	// 插入元素
	fmt.Println(li.ListInsert(4, "jjj")) // false,已满插入失败
	// 删除元素,索引为2
	li.ListDelete(2)
	// 遍历
	li.TraverseList() // {abc 10} {efg 10} GoGO

	// 根据下标Get元素
	el, _ := li.GetElem(1)
	fmt.Println(el) // {efg 10}

	// 根据传入的值，返回第一个匹配的元素下标
	b, b1 := li.LocateELem(student2)
	fmt.Println(b, b1) // 1 true

	// 寻找元素的后驱
	n1, n2 := li.NextElem(student2)
	fmt.Println(n1, n2) // GOGO true

	// 寻找元素的前驱
	p1, p2 := li.PriorElem("GOGO")
	fmt.Println(p1, p2) // {efg 10} true

	// {abc 10} {efg 10} GoGO
	// 从末尾弹出一个元素
	p1, _ = li.Pop()
	fmt.Println("从末尾弹出一个元素:", p1) // 从末尾弹出一个元素: GOGO
	li.TraverseList()             // 遍历 {abc 10} {efg 10}

	// 从末尾插入一个元素
	//b = li.Append("超哥")
	//fmt.Println(b)

	// 清空
	li.ClearList()
	fmt.Println(li.ListEmpty()) // true

	// 遍历
	li.TraverseList()
}
