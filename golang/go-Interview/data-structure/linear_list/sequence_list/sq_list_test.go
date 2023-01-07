/**
    @author:Huchao
    @data:2023/1/6
    @note:数据结构之线性表--顺序表 测试
**/
package sequence_list

import (
	"fmt"
	"testing"
)

func TestSqList(t *testing.T) {
	var li SqList
	// 初始化 1
	li.InitList(4)
	// 初始化 2
	//li = NewSeqList(4)
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

	// 扩容
	li.ExtendCapacity()
	// 获取长度
	fmt.Println("扩容后的容量：", li.Capacity)  // 8
	fmt.Println(li.ListInsert(4, "jjj")) // true
	// 遍历
	li.TraverseList() // {abc 10} {efg 10} 1000 GoGO jjj

	// 删除元素,索引为2
	li.ListDelete(2)
	// 遍历
	li.TraverseList() // {abc 10} {efg 10} GoGO jjj

	// 根据下标Get元素
	el, _ := li.GetElem(1)
	fmt.Println(el) // {efg 10}

	// 更新元素
	fmt.Println("更新元素：", li.SetElem("超哥哥", 2)) // true
	// 遍历
	li.TraverseList() // {abc 10} {efg 10} 超哥哥 jjj

	// 根据传入的值，返回第一个匹配的元素下标
	b, b1 := li.LocateELem(student2)
	fmt.Println(b, b1) // 1 true

	// 寻找元素的后驱
	n1, n2 := li.NextElem(student2)
	fmt.Println(n1, n2) // 超哥哥 true

	// 寻找元素的前驱
	p1, p2 := li.PriorElem("超哥哥")
	fmt.Println(p1, p2) // {efg 10} true

	// {abc 10} {efg 10} 超哥哥 jjj
	// 从末尾弹出一个元素
	p1 = li.Pop()
	fmt.Println("从末尾弹出一个元素:", p1) // 从末尾弹出一个元素: jjj
	li.TraverseList()             // 遍历 {abc 10} {efg 10} 超哥哥

	// 从末尾插入一个元素
	fmt.Println("从末尾插入一个元素", li.Append("超哥12")) // true
	li.TraverseList()                           // 遍历 {abc 10} {efg 10} 超哥哥 超哥12

	// 反转
	li.Reserve()
	li.TraverseList() // 遍历   超哥12 超哥哥 {efg 10} {abc 10}

	// 清空
	li.ClearList()
	fmt.Println(li.ListEmpty()) // true

	// 遍历
	li.TraverseList()
}
