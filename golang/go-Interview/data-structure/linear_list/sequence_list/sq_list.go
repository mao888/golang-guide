package sequence_list

import (
	"fmt"

	"github.com/mao888/mao-gutils/constants"
)

// 数据结构之线性表--顺序表

type SqListInterface interface {
	// 基本操作
	NewSeqList(capacity int) *SqList // 初始化
	InitList(capacity int)           // 初始化
	ListEmpty() bool                 // 判空
	ListFul() bool                   // 判满
	ListLength() int                 // 返回数据元素个数
	ClearList()                      // 清空
	DestroyList()                    // 销毁
	// 元素操作
	ListInsert(index int, elem interface{}) bool // 插入元素
	ListDelete(index int) bool                   // 删除元素
	GetElem(index int) (interface{}, bool)       // 获取元素
	SetElem(elem interface{}, index int) bool    // 更新元素
	LocateELem(elem interface{}) (int, bool)     // 返回第1个值与elem相同的元素的位置若这样的数据元素不存在,则返回值为0
	// 其他操作
	PriorElem(elem interface{}) (interface{}, bool) // 寻找元素的前驱（当前元素的前一个元素）
	NextElem(elem interface{}) (interface{}, bool)  // 寻找元素的后驱（当前元素的后一个元素）
	TraverseList()                                  // 遍历
	Pop() interface{}                               // 从末尾弹出一个元素
	Append(elem interface{}) bool                   // 从末尾插入一个元素
	ExtendCapacity()                                // 扩容
	Reserve()                                       // 反转
}

// SqList 顺序表的结构类型为SqList
// 使用golang语言的interface接口类型创建顺序表
type SqList struct {
	Len         int           // 线性表长度
	Capacity    int           // 表容量
	Data        []interface{} // 指向线性表空间
	ExtendRatio int           // 每次列表扩容的倍数
}

// NewSeqList 初始化
func (l *SqList) NewSeqList(capacity int) *SqList {
	return &SqList{
		Len:         0,
		Capacity:    capacity,
		Data:        make([]interface{}, capacity),
		ExtendRatio: constants.NumberTwo,
	}
}

// InitList 初始化
func (l *SqList) InitList(capacity int) {
	l.Capacity = capacity
	l.Len = 0
	m := make([]interface{}, capacity)
	l.Data = m
	l.ExtendRatio = constants.NumberTwo
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
		return l.Data[index], true
	}
}

// SetElem 更新元素
func (l *SqList) SetElem(elem interface{}, index int) bool {
	if index >= l.Len {
		panic("索引越界")
	}
	l.Data[index] = elem
	return true
}

// LocateELem 根据传入的值，返回第一个匹配的元素下标
func (l *SqList) LocateELem(elem interface{}) (int, bool) {
	for i, _ := range l.Data {
		if elem == l.Data[i] {
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
		pre := l.Data[i-1]
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
		N := l.Data[i+1]
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
			l.Data[i+1] = l.Data[i]
		}
		// 插入元素
		l.Data[index] = elem
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
			l.Data[i] = l.Data[i+1]
		}
		l.Len--
		return true
	}
}

// TraverseList 遍历
func (l *SqList) TraverseList() {
	for i := 0; i < l.Len; i++ {
		fmt.Println(l.Data[i])
	}
}

// ClearList 清空
func (l *SqList) ClearList() {
	l.Len = 0
	// 指针为空
	l.Data = nil
}

// DestroyList 销毁
func (l *SqList) DestroyList() {
	l.Data = nil
	l.Len = 0
	l.Capacity = 0
	l.ExtendRatio = 0
}

// Pop 从末尾弹出一个元素
func (l *SqList) Pop() interface{} {
	if l.ListEmpty() {
		panic("线性表长度为0，没有可弹出的元素")
	}
	result := l.Data[l.Len-1]
	l.Data = l.Data[:l.Len-1]
	l.Len--
	return result
}

// Append 从末尾插入一个元素
func (l *SqList) Append(elem interface{}) bool {
	if l.Len == l.Capacity {
		panic("线性表已满，无法添加数据")
	}
	l.Data = append(l.Data, elem)
	l.Len++
	return true
}

// ExtendCapacity 扩容
func (l *SqList) ExtendCapacity() {
	// 新建一个长度为 self.__size 的数组，并将原数组拷贝到新数组
	l.Data = append(l.Data, make([]interface{}, l.Capacity*(l.ExtendRatio-1))...)
	// 更新列表容量
	l.Capacity = len(l.Data)
}

// Reserve 反转
func (l *SqList) Reserve() {
	for i := 0; i < l.Len/2; i++ {
		tmp := l.Data[i]
		l.Data[i] = l.Data[l.Len-i-1]
		l.Data[l.Len-i-1] = tmp
	}
}
