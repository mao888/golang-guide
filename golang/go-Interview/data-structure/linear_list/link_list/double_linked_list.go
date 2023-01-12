package linkedlist

// 线性表 - 双向链表
// https://studygolang.com/pkgdoc

// 对链表进行遍历 (where l is a *List):
//	for e := l.Front(); e != nil; e = e.Next() {
//		// do something with e.Value
//	}
//

// Element 是链表中的元素
type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element

	// 该元素所属的链表
	list *List

	// The value stored with this element.
	Value interface{}
}

// Next 返回链表的后一个元素或者nil
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev 返回链表的前一个元素或者nil
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

type DoublyLinkedListInterface interface {
	New() *List                                        // 创建一个链表
	Init() *List                                       // 清空链表
	Len() int                                          // 返回链表中元素的个数，复杂度O(1)
	Front() *Element                                   // 返回链表第一个元素或nil
	Back() *Element                                    // 返回链表最后一个元素或nil
	PushFront(v interface{}) *Element                  // 将一个值为v的新元素插入链表的第一个位置，返回生成的新元素
	PushFrontList(other *List)                         // 创建链表other的拷贝，并将拷贝的最后一个位置连接到链表l的第一个位置
	PushBack(v interface{}) *Element                   // 将一个值为v的新元素插入链表的最后一个位置，返回生成的新元素
	PushBackList(other *List)                          // 创建链表other的拷贝，并将链表l的最后一个位置连接到拷贝的第一个位置
	InsertAfter(v interface{}, mark *Element) *Element // 将一个值为v的新元素插入到mark后面，并返回新生成的元素。如果mark不是l的元素，l不会被修改
	MoveToFront(e *Element)                            // 将元素e移动到链表的第一个位置，如果e不是l的元素，l不会被修改
	MoveToBack(e *Element)                             // 将元素e移动到链表的最后一个位置，如果e不是l的元素，l不会被修改
	MoveBefore(e, mark *Element)                       // 将元素e移动到mark的前面。如果e或mark不是l的元素，或者e==mark，l不会被修改
	MoveAfter(e, mark *Element)                        // 将元素e移动到mark的后面。如果e或mark不是l的元素，或者e==mark，l不会被修改
	Remove(e *Element) interface{}                     // 删除链表中的元素e，并返回e.Value
}

// List 代表一个双向链表。List零值为一个空的、可用的链表。
type List struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}

// Init 清空链表
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New 创建一个链表
func (l *List) New() *List {
	return new(List).Init()
}

// Len 返回链表中元素的个数，复杂度O(1)
func (l *List) Len() int {
	return l.len
}

// Front 返回链表第一个元素或nil
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back 返回链表最后一个元素或nil
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit 惰性初始化一个0 List值
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert 在at之后插入e，增加l.len，并返回e
func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// remove 从链表中删除e, l.len减1
func (l *List) remove(e *Element) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // 避免内存泄漏
	e.prev = nil // 避免内存泄漏
	e.list = nil
	l.len--
}

// move 移动e到at的next.
func (l *List) move(e, at *Element) {
	if e == at {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
}

// Remove 删除链表中的元素e，并返回e.Value。
func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// PushFront 将一个值为v的新元素插入链表的第一个位置，返回生成的新元素。
func (l *List) PushFront(v any) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack 将一个值为v的新元素插入链表的最后一个位置，返回生成的新元素。
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore 将一个值为v的新元素插入到mark前面，并返回生成的新元素。如果mark不是l的元素，l不会被修改。
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark.prev)
}

// InsertAfter 将一个值为v的新元素插入到mark后面，并返回新生成的元素。如果mark不是l的元素，l不会被修改
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List) InsertAfter(v any, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

// MoveToFront 将元素e移动到链表的第一个位置，如果e不是l的元素，l不会被修改。
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, &l.root)
}

// MoveToBack 将元素e移动到链表的最后一个位置，如果e不是l的元素，l不会被修改
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, l.root.prev)
}

// MoveBefore 将元素e移动到mark的前面。如果e或mark不是l的元素，或者e==mark，l不会被修改。
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List) MoveBefore(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter 将元素e移动到mark的后面。如果e或mark不是l的元素，或者e==mark，l不会被修改。
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List) MoveAfter(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

// PushBackList 创建链表other的拷贝，并将链表l的最后一个位置连接到拷贝的第一个位置。
// The lists l and other may be the same. They must not be nil.
func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList 创建链表other的拷贝，并将拷贝的最后一个位置连接到链表l的第一个位置
// The lists l and other may be the same. They must not be nil.
func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}
