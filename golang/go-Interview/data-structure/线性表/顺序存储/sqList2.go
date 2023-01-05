package main

import "fmt"

// https://www.jianshu.com/p/f3dedd768de4
// 数据结构之线性表--顺序表

type MyList interface {
	InitList(capacity int)                          // 初始化
	ClearList()                                     // 清空
	ListEmpty() bool                                // 判空
	ListLength()                                    // 返回数据元素个数。
	ListFul() bool                                  // 判满
	GetElem(index int) (interface{}, bool)          // 返回第i个数据元素的值
	LocateELem(elem interface{}) (int, bool)        // 返回第1个值与elem相同的元素的位置若这样的数据元素不存在,则返回值为0。
	PriorElem(elem interface{}) (interface{}, bool) // 寻找元素的前驱（当前元素的前一个元素）
	NextElem(elem interface{}) (interface{}, bool)  // 寻找元素的后驱（当前元素的后一个元素）
	ListInsert(index int, elem interface{}) bool    // 插入元素,index为插入的位置，elem为插入值
	ListDelete(index int) bool                      // 删除元素
	TraverseList()                                  // 遍历

}

// SqList 顺序表的结构类型为SqList
type SqList struct {
	Len      int            // 线性表长度
	Capacity int            // 表容量
	Prt      *[]interface{} // 指向线性表空间指针
}

// InitList 初始化
func (l *SqList) InitList(capacity int) {
	l.Capacity = capacity
	l.Len = 0
	m := make([]interface{}, capacity)
	l.Prt = &m
}

// ListEmpty 判空
func (l *SqList) ListEmpty() bool {
	if l.Len == 0 {
		return true
	} else {
		return false
	}
}

// ListLength 获取长度
func (l *SqList) ListLength() int {
	return l.Len
}

// ListFul 判满
func (l *SqList) ListFul() bool {
	if l.Len == l.Capacity {
		return true
	} else {
		return false
	}
}

// GetElem 根据下标Get元素
func (l *SqList) GetElem(index int) (interface{}, bool) {
	if index < 0 || index > l.Len {
		return nil, false
	} else {
		return (*l.Prt)[index], true
	}
}

// LocateELem 根据传入的值，返回第一个匹配的元素下标
func (l *SqList) LocateELem(elem interface{}) (int, bool) {
	for i, _ := range *l.Prt {
		if elem == (*l.Prt)[i] {
			return i, true
		}
	}
	return -1, false
}

// PriorElem 寻找元素的前驱（当前元素的前一个元素）
func (l *SqList) PriorElem(elem interface{}) (interface{}, bool) {
	i, _ := l.LocateELem(elem)
	// 顺序表中不存在该元素，或者元素为第一个元素，无前驱元素
	if i == -1 || i == 0 {
		return nil, false
	} else {
		pre := (*l.Prt)[i-1]
		return pre, true
	}
}

// NextElem 寻找元素的后驱（当前元素的后一个元素）
func (l *SqList) NextElem(elem interface{}) (interface{}, bool) {
	i, _ := l.LocateELem(elem)
	// 顺序表中不存在该元素，或者元素为最后一个元素，无后驱元素
	if i == -1 || i == l.Len-1 {
		return nil, false
	} else {
		N := (*l.Prt)[i+1]
		return N, true
	}
}

// ListInsert 插入元素,index为插入的位置，elem为插入值
func (l *SqList) ListInsert(index int, elem interface{}) bool {
	// 判断下标有效性，以及表是否满
	if index < 0 || index > l.Capacity || l.ListFul() {
		return false
	} else {
		// 先将index位置元素以及之后的元素后移一位
		for i := l.Len - 1; i >= index; i-- {
			(*l.Prt)[i+1] = (*l.Prt)[i]
		}
		// 插入元素
		(*l.Prt)[index] = elem
		l.Len++
		return true
	}
}

// DelElem 删除元素
func (l *SqList) DelElem(index int) bool {
	// 判断下标有效性，以及表是否空
	if index < 0 || index > l.Capacity || l.ListEmpty() {
		return false
	} else {
		// 注意边界
		for i := index; i < l.Len-1; i++ {
			(*l.Prt)[i] = (*l.Prt)[i+1]
		}
		l.Len--
		return true
	}
}

// TraverseList 遍历
func (l *SqList) TraverseList() {
	for i := 0; i < l.Len; i++ {
		fmt.Println((*l.Prt)[i])
	}
}

// ListClear 清空
func (l *SqList) ListClear() {
	l.Len = 0
	// 指针为空
	l.Prt = nil
}

func main() {
	var li SqList
	li.InitList(4)
	// true
	fmt.Println(li.ListEmpty())
	// false
	fmt.Println(li.ListFul())
	// 定义一个Struct类型
	type s struct {
		name string
		age  int
	}
	student1 := s{name: "abc", age: 10}
	student2 := s{name: "efg", age: 10}
	li.ListInsert(0, student1)
	li.ListInsert(1, student2)
	// false
	fmt.Println(li.ListEmpty())
	li.ListInsert(2, 1000)
	li.ListInsert(3, "GOGO")
	// {abc 10}
	// {efg 10}
	// 1000
	// GoGO
	li.TraverseList()
	// 4
	fmt.Println(li.ListLength())
	// true
	fmt.Println(li.ListFul())
	// false,已满插入失败
	fmt.Println(li.ListInsert(4, "jjj"))
	li.DelElem(2)
	// {abc 10}
	// {efg 10}
	// GoGO
	li.TraverseList()
	el, _ := li.GetElem(1)
	// {efg 10}
	fmt.Println(el)
	b, b1 := li.LocateELem(student2)
	// 1 true
	fmt.Println(b, b1)
	n1, n2 := li.NextElem(student2)
	// GOGO true
	fmt.Println(n1, n2)
	p1, p2 := li.PriorElem("GOGO")
	// {efg 10} true
	fmt.Println(p1, p2)
	li.ListClear()
	// true
	fmt.Println(li.ListEmpty())
	li.TraverseList()
}
