package my_list

import "fmt"

const (
	MAXSIZE = 20
)

type ElemType int //元素类型

const EmptyElement = 0 //空值

type MyList interface {
	//四个基本操作，初始，清空，判断是否为空，获取长度
	InitList()
	ClearList()
	ListEmpty() bool
	ListLength() int

	//四个元素操作，插入，删除，两种查找
	GetElem(index int, e *ElemType) bool
	LocateElem(value ElemType) int
	ListInsert(index int, value ElemType) bool
	ListDelete(index int, e *ElemType) bool

	Echo()
	Test()
}

//sequence_list
type SqList struct {
	Element [MAXSIZE]ElemType
	Length  int
}

//初始化列表
func (l *SqList) InitList() {
	l.Length = 0

	for i := 0; i < MAXSIZE; i++ {
		l.Element[i] = EmptyElement
	}
}

//清空列表  和初始化一个样
func (l *SqList) ClearList() {
	l.Length = 0

	for i := 0; i < MAXSIZE; i++ {
		l.Element[i] = EmptyElement
	}
}

//判断是否为空
func (l *SqList) ListEmpty() bool {
	if l.Length == 0 {
		return true
	}
	return false
}

//获取长度
func (l *SqList) ListLength() int {
	return l.Length
}

//获取指定位置的元素，返回在指针元素中
func (l *SqList) GetElem(index int, e *ElemType) bool {

	if l.Length == 0 {
		fmt.Println("获取失败，队列为空")
		return false
	}
	if index < 1 || index > l.Length {
		fmt.Println("获取失败，位置错误")
		return false
	}

	*e = l.Element[index-1]
	return true
}

//查找元素在线性表中的位置
func (l *SqList) LocateElem(value ElemType) int {

	if l.Length == 0 {
		fmt.Println("获取失败，队列为空")
		return 0
	}

	i := 0
	for ; i < MAXSIZE; i++ {
		if l.Element[i] == value {
			break
		}
	}
	if i >= MAXSIZE {
		return 0
	}
	return i + 1
}

//向线性表中指定位置插入元素
func (l *SqList) ListInsert(index int, value ElemType) bool {

	if l.Length == MAXSIZE { //满了
		fmt.Println("插入失败，队列已满")
		return false
	}
	if index < 1 || index > l.Length+1 {
		fmt.Println("插入失败，位置错误")
		return false
	}
	//开始进行插入操作，先判断元素插入位置，在将后面元素向后移动一位，进行插入
	for k := l.Length; k >= index; k-- {
		l.Element[k] = l.Element[k-1]
	}

	l.Element[index-1] = value
	l.Length++
	return true
}

//删除指定位置处的元素  并将删除的元素给e
func (l *SqList) ListDelete(index int, e *ElemType) bool {
	if l.Length == 0 {
		fmt.Println("获取失败，队列为空")
		return false
	}
	if index < 1 || index > l.Length {
		fmt.Println("获取失败，位置错误")
		return false
	}

	//删除元素，是将后面元素前移，再将最后的元素删除
	*e = l.Element[index-1]
	for k := index; k < l.Length; k++ {
		l.Element[k-1] = l.Element[k]
	}
	l.Element[l.Length-1] = EmptyElement
	l.Length--
	return true
}

//输出
func (l *SqList) Echo() {
	fmt.Println(l.Element)
}

//实现两个线性表的并集
func (la *SqList) UnionL(lb *SqList) {
	var e ElemType
	//i := 0

	La_length := la.ListLength()
	Lb_length := lb.ListLength()
	for i := 1; i <= Lb_length; i++ {
		lb.GetElem(i, &e)
		a_index := la.LocateElem(e)
		if a_index == 0 {
			La_length++
			//fmt.Println("插入index:",La_length,e)
			la.ListInsert(La_length, e)
		}
	}
}

func (l *SqList) Test() {

	fmt.Println("测试开始")

	my_list := new(SqList)
	my_list.InitList()

	for i := 1; i <= 10; i++ {
		my_list.ListInsert(i, ElemType(i*i+1))
		my_list.Echo()
	}

	fmt.Println("第5个这里插入256")
	my_list.ListInsert(5, 256)
	my_list.Echo()
	my_list.ListInsert(199, 99)

	var e ElemType

	my_list.ListDelete(1, &e)
	fmt.Println("删除头元素:", e)
	my_list.Echo()

	my_list.ListDelete(my_list.ListLength(), &e)
	fmt.Println("删除尾元素:", e)
	my_list.Echo()

	my_list.GetElem(6, &e)
	fmt.Println("获取第6个:", e)

	fmt.Println("256的位置:", my_list.LocateElem(256))

	fmt.Println("长度：", my_list.ListLength())

	fmt.Println("开始清空")
	my_list.ClearList()
	if my_list.ListEmpty() {
		fmt.Println("已清空")
		my_list.Echo()
	}

	fmt.Println("准备合并")

	my_list_a := new(SqList)
	my_list_b := new(SqList)
	for i := 1; i <= 10; i++ {
		my_list_a.ListInsert(i, ElemType(2*i+1))
		my_list_b.ListInsert(i, ElemType(3*i+1))
	}

	my_list_a.Echo()
	my_list_b.Echo()

	fmt.Println("合并后")

	my_list_a.UnionL(my_list_b)
	my_list_a.Echo()

	fmt.Println("测试完成")

}
