package main

import (
	"errors"
	"fmt"
)

// https://www.jianshu.com/p/f3dedd768de4
// 数据结构之线性表--顺序表

type MyList interface {
	InitList(capacity int)                          // 初始化
	ClearList()                                     // 清空
	ListEmpty() bool                                // 判空
	ListLength() int                                // 返回数据元素个数。
	ListFul() bool                                  // 判满
	GetElem(index int) (interface{}, bool)          // 返回第i个数据元素的值
	LocateELem(elem interface{}) (int, bool)        // 返回第1个值与elem相同的元素的位置若这样的数据元素不存在,则返回值为0。
	PriorElem(elem interface{}) (interface{}, bool) // 寻找元素的前驱（当前元素的前一个元素）
	NextElem(elem interface{}) (interface{}, bool)  // 寻找元素的后驱（当前元素的后一个元素）
	ListInsert(index int, elem interface{}) bool    // 插入元素,index为插入的位置，elem为插入值
	ListDelete(index int) bool                      // 删除元素
	TraverseList()                                  // 遍历
	Pop() (interface{}, error)                      // 从末尾弹出一个元素
	Append(elem interface{}) (bool, error)          // 从末尾插入一个元素
}

// SqList 顺序表的结构类型为SqList
// 使用golang语言的interface接口类型创建顺序表
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

// ListDelete 删除元素
func (l *SqList) ListDelete(index int) bool {
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

// ClearList 清空
func (l *SqList) ClearList() {
	l.Len = 0
	// 指针为空
	l.Prt = nil
}

// Pop 从末尾弹出一个元素
func (l *SqList) Pop() (interface{}, error) {
	if l.ListEmpty() {
		return nil, errors.New("线性表长度为0，没有可弹出的元素")
	}
	result := (*l.Prt)[l.Len-1]
	*l.Prt = (*l.Prt)[:l.Len-1]
	l.Len--
	return result, nil
}

// Append 从末尾插入一个元素
func (l *SqList) Append(elem interface{}) (bool, error) {
	if l.Len == l.Capacity {
		return false, errors.New("线性表已满，无法添加数据")
	}
	*l.Prt = append(*l.Prt, elem)
	l.Len++
	return true, nil
}
