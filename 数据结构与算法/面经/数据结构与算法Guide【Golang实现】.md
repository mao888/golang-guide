- Hello 算法：https://www.hello-algo.com/
- [https://www.topgoer.com/Go%E9%AB%98%E7%BA%A7/](https://www.topgoer.com/Go高级/)
- https://www.topgoer.cn/docs/goalgorithm/goalgorithm-1cm6akian87vb
- go算法模板：https://greyireland.gitbook.io/algorithm-pattern/
- Github: https://github.com/hunterhug

# [速查表](https://blog.csdn.net/itcodexy/article/details/109575269?app_version=5.8.0&csdn_share_tail={"type"%3A"blog"%2C"rType"%3A"article"%2C"rId"%3A"109575269"%2C"source"%3A"qq_45696377"}&utm_source=app)

## 图例

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1663837333294-8cc3292f-4ca1-4b42-a94a-cc4b44c994a5.png)

## 数据结构复杂度

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1663837000334-b5cda454-f57c-49f0-a4cc-c215ec522d12.png)

## 排序算法复杂度-数组

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1663837027297-21d925d5-00a5-4cd7-872b-366c33b243c6.png)

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671463999425-4f04dbde-e896-4575-a5a1-a9c4b93e498f.png)

## 图操作

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1663837084309-6d0dbcef-3a87-448a-937a-f511fa9e8b4e.png)

## 堆操作

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1663837277221-93bfde32-8f4a-468a-a2a1-5b80748748e4.png)

## 大-O复杂度曲线

设输入数据大小为 n ，常见的时间复杂度类型有（从低到高排列）

O(1) <O(log⁡n)<O(n)   <O(nlog⁡n)  <O(n^2) <O(2^n)<O(n!)

常数阶<对数阶<线性阶<线性对数阶<平方阶<指数阶<阶乘阶

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1672750731406-0948882c-9912-4600-8562-8b9906282079.png)

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1663837188150-fd031911-be6a-4f4f-b42c-a7a22dc61267.png)

# 数据结构

**数据结构** (Data Structure) 是相互之间存在一种或多种特定关系的数据元素的集合。换句话说，数据结构是带 ”结构＂ 的数据元素的集合， ” 就是指数据元素之间存在的关系。数据结构包括逻辑结构和存储结构两个层次。

- 整数 byte, short, int, long 、浮点数 float, double 、字符 char 、布尔 boolean 是计算机中的基本数据类型，占用空间的大小决定了它们的取值范围。
- 在程序运行时，数据存储在计算机的内存中。内存中每块空间都有独立的内存地址，程序是通过内存地址来访问数据的。
- 数据结构主要可以从逻辑结构和物理结构两个角度进行分类。逻辑结构反映了数据中元素之间的逻辑关系，物理结构反映了数据在计算机内存中的存储形式。
- 常见的逻辑结构有线性、树状、网状等。我们一般根据逻辑结构将数据结构分为线性（数组、链表、栈、队列）和非线性（树、图、堆）两种。根据实现方式的不同，哈希表可能是线性或非线性。
- 物理结构主要有两种，分别是连续空间存储（数组）和离散空间存储（链表），所有的数据结构都是由数组、或链表、或两者组合实现的。

## 分类

数据结构主要可根据「逻辑结构」和「物理结构」两种角度进行分类。

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1672839765456-7a8b8386-6161-4187-abf3-369773b2b6f1.png)

### 逻辑结构：线性与非线性

**「逻辑结构」反映了数据之间的逻辑关系。** 数组和链表的数据按照顺序依次排列，反映了数据间的线性关系；树从顶至底按层级排列，反映了祖先与后代之间的派生关系；图由结点和边组成，反映了复杂网络关系。

我们一般将逻辑结构分为「线性」和「非线性」两种。“线性”这个概念很直观，即表明数据在逻辑关系上是排成一条线的；而如果数据之间的逻辑关系是非线形的（例如是网状或树状的），那么就是非线性数据结构。

- **线性数据结构：** **线性表**（顺序表、链表）、堆栈、队列、哈希表、字符串、数组等；
- **非线性数据结构：** 树、图、堆、哈希表；

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1672752991894-12779e9b-debb-41a2-a5f6-4f2ab47c88c9.png)

### 存储结构(物理结构):顺序和链式

**「物理结构」反映了数据在计算机内存中的存储方式。** 从本质上看，分别是 **数组的连续空间存储** 和 **链表的离散空间存储** 。物理结构从底层上决定了数据的访问、更新、增删等操作方法，在时间效率和空间效率方面呈现出此消彼长的特性。

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1672753024195-c18d120e-fc31-4aaa-80e0-c292462435db.png)

**所有数据结构都是基于数组、或链表、或两者组合实现的。** 例如栈和队列，既可以使用数组实现、也可以使用链表实现，而例如哈希表，其实现同时包含了数组和链表。

- **基于数组可实现：** 栈、队列、堆、哈希表、矩阵、张量（维度 ≥3 的数组）等；
- **基于链表可实现：** 栈、队列、堆、哈希表、树、图等；

基于数组实现的数据结构也被称为「静态数据结构」，这意味着该数据结构在在被初始化后，长度不可变。相反地，基于链表实现的数据结构被称为「动态数据结构」，该数据结构在被初始化后，我们也可以在程序运行中修改其长度。

------

## 数组

### 概述

「数组 Array」是一种将 **相同类型元素** 存储在 **连续内存空间** 的数据结构，将元素在数组中的位置称为元素的「索引 Index」。

由于数组一般不做插入或删除操作， 也就是说； 一旦建立了数组， 则结构中的数据元素个数和元素之间的关系就不再发生变动。 因此， 采用**顺序存储结构**表示数组比较合适。

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1672754507162-f25890a3-f3e5-40c4-9f67-1d8bc117caf0.png)

观察上图，我们发现 **数组首元素的索引为** **0** 。你可能会想，这并不符合日常习惯，首个元素的索引为什么不是 1 呢，这不是更加自然吗？我认同你的想法，但请先记住这个设定，后面讲内存地址计算时，我会尝试解答这个问题。

**数组有多种初始化写法。** 根据实际需要，选代码最短的那一种就好。

```go
/* 初始化数组 */
var arr [5]int
nums := [5]int{1, 3, 2, 5, 4}
```

### 数组优点

**在数组中访问元素非常高效。** 这是因为在数组中，计算元素的内存地址非常容易。给定数组首个元素的地址、和一个元素的索引，利用以下公式可以直接计算得到该元素的内存地址，从而直接访问此元素。

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1672754625641-694830e7-57c8-4c10-a2cd-fe96f76d5cc8.png)

```go
// 元素内存地址 = 数组内存地址 + 元素长度 * 元素索引
elementAddr = firtstElementAddr + elementLength * elementIndex
```

**为什么数组元素索引从 0 开始编号？** 根据地址计算公式，**索引本质上表示的是内存地址偏移量**，首个元素的地址偏移量是 0 ，那么索引是 0 也就很自然了。

访问元素的高效性带来了许多便利。例如，我们可以在 O(1) 时间内随机获取一个数组中的元素。

```go
/* 随机返回一个数组元素 */
func randomAccess(nums []int) (randomNum int) {
    // 在区间 [0, nums.length) 中随机抽取一个数字
    randomIndex := rand.Intn(len(nums))
    // 获取并返回随机元素
    randomNum = nums[randomIndex]
    return
}
```

### 数组缺点

**1、数组在初始化后长度不可变。** 由于系统无法保证数组之后的内存空间是可用的，因此数组长度无法扩展。而若希望扩容数组，则需新建一个数组，然后把原数组元素依次拷贝到新数组，在数组很大的情况下，这是非常耗时的。

**2、数组中插入或删除元素效率低下。** 假设我们想要在数组中间某位置插入一个元素，由于数组元素在内存中是“紧挨着的”，它们之间没有空间再放任何数据。因此，我们不得不将此索引之后的所有元素都向后移动一位，然后再把元素赋值给该索引。删除元素也是类似，需要把此索引之后的元素都向前移动一位。总体看有以下缺点：

- **时间复杂度高：** 数组的插入和删除的平均时间复杂度均为 O(N) ，其中 N 为数组长度。
- **丢失元素：** 由于数组的长度不可变，因此在插入元素后，超出数组长度范围的元素会被丢失。
- **内存浪费：** 我们一般会初始化一个比较长的数组，只用前面一部分，这样在插入数据时，丢失的末尾元素都是我们不关心的，但这样做同时也会造成内存空间的浪费。

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1672755014104-2eece5f7-df6d-4334-8ae0-99a3a6f9cae7.png)

### 数组常用操作

```go
package 数组与链表

import (
	"math/rand"
)

/**
我们将 Go 中的 Slice 切片看作 Array 数组，降低理解成本，
有利于我们将关注点放在数据结构与算法上。
*/

/* 随机返回一个数组元素 */
func randomAccess(nums []int) (randomNum int) {
	// 在区间 [0, nums.length) 中随机抽取一个数字
	randomIndex := rand.Intn(len(nums))
	// 获取并返回随机元素
	randomNum = nums[randomIndex]
	return
}

/* 扩展数组长度 */
func extend(nums []int, enlarge int) []int {
	// 初始化一个扩展长度后的数组
	res := make([]int, len(nums)+enlarge)
	// 将原数组中的所有元素复制到新数组
	for i, num := range nums {
		res[i] = num
	}
	// 返回扩展后的新数组
	return res
}

/* 在数组的索引 index 处插入元素 num */
func insert(nums []int, num int, index int) {
	// 把索引 index 以及之后的所有元素向后移动一位
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}
	// 将 num 赋给 index 处元素
	nums[index] = num
}

/* 删除索引 index 处元素 */
func remove(nums []int, index int) {
	// 把索引 index 之后的所有元素向前移动一位
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}

/* 遍历数组 */
func traverse(nums []int) {
	count := 0
	// 通过索引遍历数组
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[count])
		count++
	}
	count = 0
	// 直接遍历数组
	for range nums {
		fmt.Println(nums[count])
		count++
	}
}

/* 在数组中查找指定元素 */
func find(nums []int, target int) (index int) {
	index = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			index = i
			break
		}
	}
	return
}
package array

/**
我们将 Go 中的 Slice 切片看作 Array array。因为这样可以
降低理解成本，利于我们将关注点放在数据结构与算法上。
*/

import (
	"fmt"
	"testing"
)

/* Driver Code */
func TestArray(t *testing.T) {
	/* 初始化数组 */
	var arr []int
	fmt.Println("array arr =", arr)
	nums := []int{1, 3, 2, 5, 4}
	fmt.Println("array nums =", nums)

	/* 随机访问 */
	randomNum := randomAccess(nums)
	fmt.Println("在 nums 中获取随机元素", randomNum)

	/* 长度扩展 */
	nums = extend(nums, 3)
	fmt.Println("将数组长度扩展至 8 ，得到 nums =", nums)

	/* 插入元素 */
	insert(nums, 6, 3)
	fmt.Println("在索引 3 处插入数字 6 ，得到 nums =", nums)

	/* 删除元素 */
	remove(nums, 2)
	fmt.Println("删除索引 2 处的元素，得到 nums =", nums)

	/* 遍历数组 */
	traverse(nums)

	/* 查找元素 */
	index := find(nums, 3)
	fmt.Println("在 nums 中查找元素 3 ，得到索引 =", index)
}
```

### 数组典型应用

1. **随机访问。** 如果我们想要随机抽取一些样本，那么可以用数组存储，并生成一个随机序列，根据索引实现样本的随机抽取。
2. **二分查找。** 例如前文查字典的例子，我们可以将字典中的所有字按照拼音顺序存储在数组中，然后使用与日常查纸质字典相同的“翻开中间，排除一半”的方式，来实现一个查电子字典的算法。
3. **深度学习。** 神经网络中大量使用了向量、矩阵、张量之间的线性代数运算，这些数据都是以数组的形式构建的。数组是神经网络编程中最常使用的数据结构。
    

------

## 线性表

### 基本概念

- `**定义**：零个或者多个数据元素的有限序列，在复杂的线性表中，一个数据元素可以由若干个数据项组成。`
- `直接前驱元素：若线性表记为(a1a2a3...an),则表中a2领先于a3，则称a2是a3的直接前驱元素，且有且仅有一个直接前驱元素`
- `直接后继元素：称a3是a2的直接后继元素，且有且仅有一个直接后继元素`
- `线性表的长度：线性表的元素个数n，为线性表的长度，随着线性表插入和删除操作，该值是变动的，线性表的存储长度一般小于数组的长度`
- `数组的长度：存放线性表的存储空间的长度，存储分配后这个值一般是不变的`
- `空表：长度n为0时，该线性表为空表`
- `地址：存储器的每个存储单元都有自己在内存的编号，简称为地址`

#### 线性结构的定义

若结构是非空有限集，则有且仅有一个开始结点和一个终端结点，并且所有结点都最多只有一个直接前趋和一个直接后继。

#### 线性结构的特点

- ① 只有一个首结点和尾结点；
- ② 除首尾结点外，其他结点只有一个直接前驱和一个直接后继。
- 简言之，线性结构反映结点间的逻辑关系是 一对一  的

线性结构包括 线性表、堆栈、队列、字符串、数组等等，其中，最典型、最常用的是：**线性表**

### 顺序表

#### 基本概述

- 线性表的顺序表示又称为顺序存储结构或顺序映像，是线性表的一种。
- 顺序表不仅要求数据在逻辑上是连续的一条直线，还要求用一段**物理地址连续的存储单元**以次存储表中数据元素，一般情况下采用数组存储。



- **顺序存储定义：**把逻辑上相邻的数据元素存储在物理上相邻的存储单元中的存储结构。
- 简言之，逻辑上相邻，物理上也相邻
- **顺序存储方法：**用一组地址连续的存储单元依次存储线性表的元素，可通过数组V[n]来实现。

#### 基本操作的代码实现

```go
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
	Reserve()                                       // 反转
}

// SqList 顺序表的结构类型为SqList
// 使用golang语言的interface接口类型创建顺序表
type SqList struct {
	Len         int           // 线性表长度
	Capacity    int           // 表容量
	Data        []interface{} // 指向线性表空间指针
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
```

#### 测试代码

```go
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
```

#### 复杂度

**取值** **O(1)**

只要i 的数值在数组下标范围内，就是把数组第 i - 1 下标的值返回即可，顺序表取值算法的时间复杂度为：**O(1)**

**查找** **O(n)**

顺序表按值查找算法的平均时间复杂度为 **O(n)**

**插入** **O(n)**



1. 若插入在尾结点之后，则根本无需移动（特别快）；
2. 若插入在首结点之前，则表中元素全部后移（特别慢）；
3. 在各种位置插入（共n+1种可能）的平均移动次数

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1673021483253-fb6c4189-307b-444b-ae91-1df22fbc856b.png)

由此可见， 顺序表插入算法的平均时间复杂度为 **O(n)**。

**删除** **O(n)**

1. 若删除尾结点，则根本无需移动（特别快）；
2. 若删除首结点，则表中n-1个元素全部前移（特别慢）；
3. 若要考虑在各种位置删除（共n种可能）的平均移动次数

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1673021587107-3dba8712-ad36-4e7d-a59f-9fe020a71880.png)

顺序表删除算法的平均时间复杂度为：**O(n)**。

顺序表可以随机存取表中任一元素，其存储位置可用一个简单、直观的公式来表示。然而，从另一方面来看，这个特点也造成了这种存储结构的缺点：在做插入或删除操作时，需移动大最元素。 另外由于数组有长度相对固定的静态特性， 当表中数据元素个数较多且变化较大时，操作过程相对复杂，必然导致存储空间的浪费。 所有这些问题，都可以通过线性表的另一种表示方法——链式存储结构来解决。

#### 顺序表（顺序存储结构）的特点

（1）利用数据元素的存储位置表示线性表中相邻数据元素之间的前后关系，即线性表的逻辑结构与存储结构一致

（2）在访问线性表时，可以快速地计算出任何一个数据元素的存储地址。因此可以粗略地认为，访问每个元素所花时间相等　

#### 顺序表的优缺点 

**优点**

1. **存储密度大**（结点本身所占存储量/结点结构所占存储量）
2. 无须为表示表中元素之间的逻辑关系而增加额外的存储空间
3. 可以快速地存取表中任一位置的元素

**缺点**

1. 在插入、删除某一元素时，需要移动大量元素
2. 当线性表长度变化较大肘，难以确定存储空间的容量，造成存储空间的"碎片"，浪费存储
3. 空间属于静态存储形式，数据元素的个数不能自由扩充

为克服这一缺点 ==》 链表

### 链表

##### 链式存储结构



- `特点：是用一组任意的存储单元存储线性表的数据元素，可以是连续的也可以是不连续的。`
- `数据域：为了表示每个数据元素ai与其直接后继元素ai+1之间的逻辑关系，对数据元素ai来说，除了存储其本身的信息之外，还需要存储一个指示其直接后继的信息（即直接后继的存储位置)，存储信息的域叫数据域`
- `指针域：把存储直接后继位置的域称为指针域`
- `指针|链：指针域中存储的信息称为指针或域`
- `结点：数据域和指针域组成数据元素ai的存储映像，称为结点`
- `头指针：把链表中第一个结点的存储位置叫做头指针，线性表的最后一个结点指针为空`
- `头结点：在单链表的第一个结点前附设一个结点，称为头结点，头结点的数据域可以不存储任何信息，也可以存储线性表的长度等附加信息，头结点的指针域存储指向第一个结点的指针。`



###### 单链表



```plain
单链表：n个结点链接成一个链表，即为线性表(a1a2a3...an)的链式存储结构，因为此链表的每个结点中只包含一个指针域，所以叫做单链表。单链表正是通过每个结点的指针域将线性表的数据元素按其逻辑次序链接在一起的。
```



```go
package linelist

// 单链表结点
type SingleList struct {
    Data interface{} //单链表的数据域
    Next *SingleList //单链表的指针域
}

func NewSingleList() *SingleList {
    return &SingleList{Data: "", Next: nil}
}

type SingleListr interface {
    GetFirst() *SingleList
    GetLast() *SingleList
    Length() int
    Add(data interface{}) bool
    GetElem(index int) (interface{}, error)
    Delete(index int) bool
}

//返回第一个结点
func (this *SingleList) GetFirst() *SingleList {
    if this.Next == nil {
        return nil
    }
    return this.Next
}

//返回最后一个结点
func (this *SingleList) GetLast() *SingleList {
    if this.Next == nil {
        return nil
    }
    point := this
    for point.Next != nil {
        point = point.Next
    }
    if point.Next == nil {
        return point
    }
    return nil
}

//获取单链表的长度
func (this *SingleList) Length() int {
    point := this
    length := 0

    for point.Next != nil {
        length++
        point = point.Next
    }
    return length
}

//往单链表的末尾加一个元素
func (this *SingleList) Add(data interface{}) bool {
    point := this
    for point.Next != nil {
        point = point.Next
    }
    tmpSingle := SingleList{Data: data}
    point.Next = &tmpSingle
    return true
}

//获取所有结点的值
func (this *SingleList) GetAll() []interface{} {
    result := make([]interface{}, 0)
    point := this
    for point.Next != nil {
        result = append(result, point.Data)
        point = point.Next
    }
    result = append(result, point.Data)
    return result
}

//获取索引为index的结点
func (this *SingleList) GetElem(index int) *SingleList {
    point := this
    if index < 0 || index > this.Length() {
        panic("check index error")
        return nil
    }
    for i := 0; i < index; i++ {
        point = point.Next
    }
    return point
}

//删除第index个结点
func (this *SingleList) Delete(index int) bool {
    if index < 0 || index > this.Length() {
        panic("please check index")
        return false
    }
    point := this
    for i := 0; i < index-1; i++ {
        point = point.Next
    }
    point.Next = point.Next.Next
    return true
}
```



###### 单循环链表



```plain
定义：将单链表中终端节点的指针端由空指针改为指向头节点，就使得整个单链表形成一个环，这种头尾相接的单链表简称为循环链表
```



```go
package linelist

import "errors"

//定义单循环链表的节点数据结构
type CircleNode struct {
    data interface{}
    next *CircleNode
}

//定义单循环链表的数据结构
type CircleList struct {
    tail *CircleNode
    size int
}

func InitCircleList() *CircleList {
    return &CircleList{tail: nil, size: 0}
}

func InitCircleNode(data interface{}) *CircleNode {
    return &CircleNode{data: data, next: nil}
}

//单链表在表尾添加数据
func (cl *CircleList) Append(data *CircleNode) bool {
    if data == nil {
        return false
    }
    if cl.size == 0 {
        data.next = data
    } else {
        curNode := cl.tail.next
        data.next = curNode
        cl.tail.next = data
    }
    cl.tail = data
    cl.size++
    return true
}

//单循环链表插入数据
func (cl *CircleList) Insert(num int, data *CircleNode) error {
    if data == nil {
        return errors.New("要插入的节点数据为空")
    }
    if cl.size == 0 || cl.size == num {
        cl.Append(data)
    } else {
        var curNode *CircleNode
        if num == 0 {
            curNode = cl.tail
        } else {
            curNode = cl.Get(num)
            if cl.size == num {
                cl.tail = data
            }
        }
        data.next = curNode.next
        curNode.next = data
        cl.size++
    }
    return nil
}

//单循环链表查询数据
func (cl *CircleList) Get(num int) *CircleNode {
    if num < 0 || num > cl.size-1 {
        return nil
    }
    curNode := cl.tail
    for i := 0; i < num; i++ {
        curNode = curNode.next
    }
    return curNode
}

//单循环链表查询全部数据
func (cl *CircleList) GetAll() []interface{} {
    result := make([]interface{}, 0)
    curNode := cl.tail
    for i := 0; i < cl.size; i++ {
        result = append(result, curNode.data)
        curNode = curNode.next
    }
    return result
}

//单循环链表按序号删除数据
func (cl *CircleList) RemoveInt(num int) error {
    if cl.size == 0 {
        return errors.New("循环链表为空")
    }
    if num > cl.size-1 {
        return errors.New("越界")
    }

    if cl.size == 1 {
        cl.tail = nil
        cl.size = 0
        return nil
    } else {
        var curNode *CircleNode
        var data *CircleNode
        if num == 0 {
            curNode = cl.tail
        } else {
            curNode = cl.Get(num - 1)
        }

        data = curNode.next
        curNode.next = data.next

        if num == cl.size-1 {
            cl.tail = curNode
        }

        data.next = nil
        data = nil
        cl.size--

        return nil
    }
}

//单循环链表删除全部数据
func (cl *CircleList) RemoveAll() bool {
    if cl.size == 0 {
        return false
    }

    for i := 0; i < cl.size; i++ {
        curNode := cl.tail
        cl.tail = curNode.next
        curNode.next = nil
    }
    cl.tail = nil
    cl.size = 0

    return true
}
```



###### 双向链表



```plain
定义：在单链表的每个节点中，再设置一个指向其前驱节点的指针域。所以在双向链表中的节点都有两个指针域，一个指向直接后继，另一个直接指向前驱
```



```go
package linelist

import (
    "errors"
)

var (
    NUMERROR = errors.New("链表越界")
)
//定义双向链表节点结构体
type DoubleNode struct {
    data interface{}
    prev *DoubleNode
    next *DoubleNode
}

//定义双向链表结构体
type DoubleList struct {
    head *DoubleNode
    tail *DoubleNode
    size int
}

//初始化链表

func InitDoubleList() *DoubleList {
    return &DoubleList{head: nil, tail: nil, size: 0}
}

func InitDoubleNode(data interface{}) *DoubleNode {
    return &DoubleNode{data: data, prev: nil, next: nil}
}

//获取链表的长度
func (dl *DoubleList) GetSize() int {
    return dl.size
}

//获取链表头部节点
func (dl *DoubleList) GetHead() *DoubleNode {
    return dl.head
}

//获取链表尾部节点
func (dl *DoubleList) GetTail() *DoubleNode {
    return dl.tail
}

//在头部追加节点
func (dl *DoubleList) AddHeadNode(node *DoubleNode) int {
    if dl.GetSize() == 0 {
        dl.head = node
        dl.tail = node
        node.prev = nil
        node.next = nil
    } else {
        dl.head.prev = node
        node.prev = nil
        node.next = dl.head
        dl.head = node
    }
    dl.size += 1
    return dl.size
}

//在尾部追加节点
func (dl *DoubleList) AddTailNode(node *DoubleNode) int {
    if dl.GetSize() == 0 {
        dl.head = node
        dl.tail = node
        node.prev = nil
        node.next = nil
    } else {
        dl.tail.next = node
        node.prev = dl.tail
        node.next = nil
        dl.tail = node
    }
    dl.size += 1
    return dl.size
}

//在链表某个序号之后插入节点
func (dl *DoubleList) InsertNextInt(num int, data *DoubleNode) bool {
    if data == nil || num > dl.GetSize()-1 || num < 0 {
        return false
    }
    switch {
    case dl.GetSize() == 0:
        dl.AddHeadNode(data)
    case num == dl.GetSize()-1:
        dl.AddTailNode(data)
    default:
        curNode, err := dl.GetOrder(num)
        if err != nil {
            return false
        }
        data.prev = curNode
        data.next = curNode.next
        curNode.next = data
        curNode.next.prev = data
        dl.size++
    }
    return true
}

//顺序查询某个序号的数据
func (dl *DoubleList) GetOrder(num int) (*DoubleNode, error) {
    switch {
    case dl.GetSize() == 0:
        return nil, NUMERROR
    case num == 0:
        return dl.head, nil
    case num > dl.GetSize()-1:
        return nil, NUMERROR
    case num == dl.GetSize()-1:
        return dl.tail, nil
    default:
        data := dl.head
        for i := 0; i < num; i++ {
            data = data.next
        }
        return data, nil
    }
}

//倒序查询某个序号数据
func (dl *DoubleList) GetReverse(num int) (data *DoubleNode, err error) {
    switch {
    case num == 0:
        data = dl.tail
    case num > dl.GetSize()-1:
        err = NUMERROR
    case num == dl.GetSize()-1:
        data = dl.head
    default:
        data = dl.tail
        for i := 0; i < num; i++ {
            data = data.prev
        }
    }
    return
}

//获取链表中所有数据
func (dl *DoubleList) GetAll() []interface{} {
    result := make([]interface{}, 0)
    if dl.GetSize() == 0 {
        return nil
    }
    curNode := dl.head
    for i := 0; i < dl.GetSize(); i++ {
        result = append(result, curNode.data)
        curNode = curNode.next
    }
    return result
}

//删除某个序号的数据
func (dl *DoubleList) Remove(num int) error {
    if dl.GetSize() == 0 {
        return NUMERROR
    }

    var curNode *DoubleNode
    var err error
    if curNode, err = dl.GetOrder(num); err != nil {
        return err
    }

    if num == 0 {
        curNode.next.prev = nil
        dl.head = curNode.next
    } else if num == dl.size-1 {
        curNode.prev.next = nil
        dl.tail = curNode.prev
    } else {
        curNode.prev.next = curNode.next
        curNode.next.prev = curNode.prev
    }

    curNode.prev = nil
    curNode.next = nil
    dl.size--
    return nil
}

//删除链表中的全部数据
func (dl *DoubleList) RemoveAll() bool {
    for i := 0; i < dl.GetSize(); i++ {
        curNode := dl.head
        dl.head = curNode.next
        curNode.next = nil
        curNode.prev = nil
    }
    dl.tail = nil
    dl.size = 0
    return true
}
```

------

## 数组、线性表的区别

**线性表**

1. 线性表（linear list）是数据结构的一种，一个线性表是 n 个具有相同特性的数据元素的有限序列。
2. 线性表在逻辑上是线性结构，也就说是连续的一条直线，但是在物理结构上并不一定是连续的。
3. 常见的线性表：顺序表、链表、栈、队列、字符串

**顺序表**

1. 顺序表，全名顺序存储结构，是线性表的一种。
2. 顺序表不仅要求数据在逻辑上是连续的一条直线，还要求用一段物理地址连续的存储单元以存储表中数据元素，一般情况下采用数组存储。

**什么是数组**

1. 数组是相同数据类型的元素按一定顺序排列的的集合。数组中的元素存储在一个连续性的内存块中，并通过索引来访问。
2. 简单的说，数组是在物理空间中连续存储的相同数据类型的元素的集合。

**总结**

1. **数组**是数据结构中顺序存储的物理结构，而**顺序表**是数据结构中的逻辑结构
2. **顺序表**是从逻辑结构的角度来说的，它的每一个元素都只有一个前驱元素和一个后驱元素除了头和尾，逻辑结构还有队列，堆栈，树，图等
3. **数组**是从物理存贮的角度来说的，顺序表用数组存贮也可以用链表来存贮。
4. 可以用**数组**实现顺序表，但我们同样可以用数组实现二叉树、队列等结构，因此不能直接认为顺序表就是数组

## Q1：链表，队列和栈的区别 

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1662348730108-fb8ddbb0-b395-4e57-802c-e85476e5146e.png)



- **链表**是一种物理存储单元上非连续的一种数据结构，看名字我们就知道他是一种链式的结构，就像一群人手牵着手一样。链表有单向的，双向的，还有环形的。
- **队列**是一种特殊的线性表，他的特殊性在于我们只能操作他头部和尾部的元素，中间的元素我们操作不了，我们只能在他的头部进行删除，尾部进行添加。就像大家排队到银行取钱一样，先来的肯定要排到前面，后来的只能排在队尾，所有元素都要遵守这个操作，没有VIP会员，所以走后门插队的现象是不可能存在的，他是一种先进先出的数据结构。我们来看一下队列的数据结构是什么样的。
- **栈**也是一种特殊的线性表，他只能对栈顶进行添加和删除元素。栈有入栈和出栈两种操作，他就好像我们把书一本本的摞起来，最先放的书肯定是摞在下边，最后放的书肯定是摞在了最上面，摞的时候不允许从中间放进去，拿书的时候也是先从最上面开始拿，不允许从下边或中间抽出来。

## Q2:二叉树中完全二叉树、满二叉树、二叉排序树、平衡二叉树的区别和联系

### 1、完全二叉树：

只有最下面的两层结点度小于2，并且最下面一层的结点都集中在该层最左边的若干位置。

### 2、满二叉树：

是一颗完全二叉树；

除了叶结点外每一个结点都有左右子叶且叶结点都处在最底层。深度为k，且有2的(k)次方－1个节点。

### 3、堆：

是一颗完全二叉树；

大根堆：左右子树的结点值都小于根结点值，左右子树都是大根堆。

小根堆：左右子树的结点值都大于根结点值，左右子树都是小根堆。

### 4、[二叉排序树（二叉查找树）](https://www.topgoer.com/Go高级/二叉搜索树.html)及代码实现:

左子树上的值都小于根结点的值，右子树上的值都大于根结点得值，左右子树都是二叉排序树。

### 5、平衡二叉树（ALV）：

是一颗二叉排序树；

左子树和右子树的差值不超过1，左右子树都为平衡二叉树。

常用算法有红黑树、AVL、Treap

## Q5：什么是 AVL 树？

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1662348730387-6e2c3794-187b-4c7d-8f97-27850a7bbe25.png)



**AVL 树** 是平衡⼆叉查找树，增加和删除节点后通过树形旋转重新达到平衡。右旋是以某个节点为中⼼， 将它沉⼊当前右⼦节点的位置，⽽让当前的左⼦节点作为新树的根节点，也称为顺时针旋转。同理左旋 是以某个节点为中⼼，将它沉⼊当前左⼦节点的位置，⽽让当前的右⼦节点作为新树的根节点，也称为 逆时针旋转。



## Q6：什么是红⿊树？



红⿊树 是 1972 年发明的，称为对称⼆叉 B 树，1978 年正式命名红⿊树。主要特征是在每个节点上增加⼀个属性表示节点颜⾊，可以红⾊或⿊⾊。





![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1662348730546-f3de5e10-20eb-44b5-823f-e5237bd974c4.png)



红⿊树和 AVL 树 类似，都是在进⾏插⼊和删除时通过旋转保持⾃身平衡，从⽽获得较⾼的查找性能。与 AVL 树 相⽐，红⿊树不追求所有递归⼦树的⾼度差不超过 1，保证从根节点到叶尾的最⻓路径不超过最短路径的 2 倍，所以最差时间复杂度是 O(logn)。

红⿊树通过重新着⾊和左右旋转，更加⾼效地完成了插⼊和删除之后的⾃平衡调整。红⿊树在本质上还是⼆叉查找树，它额外引⼊了 5 个约束条件： ① 节点只能是红⾊或⿊⾊。 ② 根节点必须是⿊⾊。 ③ 所有 NIL 节点都是⿊⾊的。 ④ ⼀条路径上不能出现相邻的两个红⾊节点。 ⑤ 在任何递归⼦树中，根节点到叶⼦节点的所有路径上包含相同数⽬的⿊⾊节点。

这五个约束条件保证了红⿊树的新增、删除、查找的最坏时间复杂度均为 O(logn)。如果⼀个树的左⼦节点或右⼦节点不存在，则均认定为⿊⾊。红⿊树的任何旋转在 3 次之内均可完成。



## Q7：AVL 树和红⿊树的区别？

红⿊树的平衡性不如 AVL 树，它维持的只是⼀种⼤致的平衡，不严格保证左右⼦树的⾼度差不超过 1。这导致节点数相同的情况下，红⿊树的⾼度可能更⾼，也就是说平均查找次数会⾼于相同情况的 AVL 树。

在插⼊时，红⿊树和 AVL 树都能在⾄多两次旋转内恢复平衡，在删除时由于红⿊树只追求⼤致平衡，因此红⿊树⾄多三次旋转可以恢复平衡，⽽ AVL 树最多需要 O(logn) 次。AVL 树在插⼊和删除时，将向上回溯确定是否需要旋转，这个回溯的时间成本最差为 O(logn)，⽽红⿊树每次向上回溯的步⻓为 2，回溯成本低。因此⾯对频繁地插⼊与删除红⿊树更加合适。

## Q8：B 树和B+ 树的区别？





![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1662348730557-7fb979a5-1145-42f7-9bbf-3115190b4e24.png)



B 树中每个节点同时存储 key 和 data，⽽ B+ 树中只有叶⼦节点才存储 data，⾮叶⼦节点只存储 key。InnoDB 对 B+ 树进⾏了优化，在每个叶⼦节点上增加了⼀个指向相邻叶⼦节点的链表指针，形成了带有顺序指针的 B+ 树，提⾼区间访问的性能。

B+ 树的优点在于： ① 由于 B+ 树在⾮叶⼦节点上不含数据信息，因此在内存⻚中能够存放更多的key，数据存放得更加紧密，具有更好的空间利⽤率，访问叶⼦节点上关联的数据也具有更好的缓存命 中率。 ② B+树的叶⼦结点都是相连的，因此对整棵树的遍历只需要⼀次线性遍历叶⼦节点即可。⽽ B 树则需要进⾏每⼀层的递归遍历，相邻的元素可能在内存中不相邻，所以缓存命中性没有 B+树好。但是 B 树也有优点，由于每个节点都包含 key 和 value，因此经常访问的元素可能离根节点更近，访问也更迅速。



## [Q9：图](https://www.topgoer.com/Go高级/图.html)



## [Q10：散列表](https://www.topgoer.com/Go高级/散列表.html)



## [Q11：堆](https://www.topgoer.com/Go高级/堆.html)



## [Q12：链表](https://www.topgoer.com/Go高级/链表.html)



## [Q13：跳跃表](https://www.topgoer.com/Go高级/跳跃表.html)



## [Q14：字典树](https://www.topgoer.com/Go高级/字典树.html)



## [Q15：向量空间](https://www.topgoer.com/Go高级/向量空间.html)

# 排序算法

## 排序有哪些分类？

### 1)内部排序

指将需要处理的所有数据都加载到内部存储器(内存)中进行排序。

内部排序包括⽐较排序和⾮⽐较排序，⽐较排序包括插⼊/选择/交换/归并排序，⾮⽐较排序包括计数/ 基数/桶排序。

插⼊排序包括直接插⼊/希尔排序，选择排序包括直接选择/堆排序，交换排序包括冒泡/快速排序。

### 2)外部排序法

数据量过大，无法全部加载到内存中，需要借助外部存储(文件等)进行排序。

### 3）常见的排序算法分类(如图)

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671462941318-12f62558-3b48-4b4d-abb9-adfb806c4428.png)



## 时间、空间复杂度、稳定性对比

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671463991125-53080e1c-99d5-4024-8fbb-1e7f44110614.png)

**稳定**：如果a原本在b前面，而a=b，排序之后a仍然在b的前面。

**不稳定**：如果a原本在b的前面，而a=b，排序之后 a 可能会出现在 b 的后面。

### 时间复杂度：

- 「时间复杂度」统计算法运行时间随着数据量变大时的增长趋势，可以有效评估算法效率，但在某些情况下可能失效，比如在输入数据量较小或时间复杂度相同时，无法精确对比算法效率的优劣性。
- 「最差时间复杂度」使用大 O 符号表示，即函数渐近上界，其反映当 n 趋于正无穷时，T(n) 处于何种增长级别。
- 推算时间复杂度分为两步，首先统计计算操作数量，再判断渐近上界。
- 常见时间复杂度从小到大排列有 O(1) , O(log⁡n) , O(n) , O(nlog⁡n) , O(n^2) , O(2^n) , O(n!) 。
- 某些算法的时间复杂度不是恒定的，而是与输入数据的分布有关。时间复杂度分为「最差时间复杂度」和「最佳时间复杂度」，后者几乎不用，因为输入数据需要满足苛刻的条件才能达到最佳情况。
- 「平均时间复杂度」可以反映在随机数据输入下的算法效率，最贴合实际使用情况下的算法性能。计算平均时间复杂度需要统计输入数据的分布，以及综合后的数学期望。

1.时间复杂度

一般情况下，算法中基本操作重复执行的次数是问题规模n的某个函数，用T(n)表示，若有某个辅助函数f(n),使得当n趋近于无穷大时，T（n)/f(n)的极限值为不等于零的常数，则称f(n)是T(n)的同数量级函数。记作T(n)=Ｏ(f(n)),称Ｏ(f(n)) 为算法的渐进时间复杂度，简称时间复杂度。

上面这一段解释是很规范的，但是对于非专业性的我们来说并不是那么好理解，说白了时间复杂度就是时间复杂度的计算并不是计算程序具体运行的时间，而是算法执行语句的次数。通常我们计算时间复杂度都是计算最坏情况 。

#### 最好时间复杂度:

在完全有序的情况下的时间复杂度（满有序度）如(1，2，3)

#### 最坏时间复杂度

最好时间复杂度:在完全有序的情况下的时间复杂度（满有序度)如（1，2，3)

#### 平均时间复杂度

平均时间复杂度是指所有可能的输入实例均以等概率出现的情况下，算法的期望运行时间。设每种情况的出现的概率为pi,平均时间复杂度则为sum(pi*f(n))

### 空间复杂度：

- 与时间复杂度的定义类似，「空间复杂度」统计算法占用空间随着数据量变大时的增长趋势。
- 算法运行中相关内存空间可分为输入空间、暂存空间、输出空间。通常情况下，输入空间不计入空间复杂度计算。暂存空间可分为指令空间、数据空间、栈帧空间，其中栈帧空间一般在递归函数中才会影响到空间复杂度。
- 我们一般只关心「最差空间复杂度」，即统计算法在「最差输入数据」和「最差运行时间点」下的空间复杂度。
- 常见空间复杂度从小到大排列有 O(1) , O(log⁡n) , O(n) , O(n^2) , O(2^n) 。

一个程序的空间复杂度是指运行完一个程序所需**内存的大小**。利用程序的空间复杂度，可以对程序的运行所需要的内存多少有个预先估计。一个程序执行时除了需要存储空间和存储本身所使用的指令、常数、变量和输入数据外，还需要一些对数据进行操作的工作单元和存储一些为现实计算所需信息的辅助空间。程序执行时所需存储空间包括以下两部分。　　

（1）**固定部分**。这部分空间的大小与输入/输出的数据的个数多少、数值无关。主要包括指令空间（即代码空间）、数据空间（常量、简单变量）等所占的空间。这部分属于静态空间。

（2）**可变空间**，这部分空间的主要包括动态分配的空间，以及递归栈所需的空间等。这部分的空间大小与算法有关。

一个算法所需的存储空间用f(n)表示。S(n)=O(f(n))　　其中n为问题的规模，S(n)表示空间复杂度。 

### 稳定性：

如果待排序的序列中存在值相等的元素，经过排序之后，相等元素之间原有的先后顺序不变

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671463906873-1d75eacb-da8f-43f7-92d0-b4d474d965e6.png)

稳定也可以理解为一切皆在掌握中,元素的位置处在你在控制中.而不稳定算法有时就有点碰运气,随机的成分.当两元素相等时它们的位置在排序后可能仍然相同.但也可能不同.是未可知的.

另外要注意的是:算法思想的本身是独立于编程语言的,所以你写代码去实现算法的时候很多细节可以做不同的处理.采用不稳定算法不管你具体实现时怎么写代码,最终相同元素位置总是不确定的(可能位置没变也可能变了).而稳定排序算法是你在具体实现时如果细节方面处理的好就会是稳定的,但有些细节没处理得到的结果仍然是不稳定的.

比如冒泡排序,直接插入排序,归并排序虽然是稳定排序算法,但如果你实现时细节没处理好得出的结果也是不稳定的.

### 稳定性的用处

我们平时自己在使用排序算法时用的测试数据就是简单的一些数值本身.没有任何关联信息.这在实际应用中一般没太多用处.实际应该中肯定是排序的数值关联到了其他信息,比如数据库中一个表的主键排序,主键是有关联到其他信息.另外比如对英语字母排序,英语字母的数值关联到了字母这个有意义的信息.

可能大部分时候我们不用考虑算法的稳定性.两个元素相等位置是前是后不重要.但有些时候稳定性确实有用处.它体现了程序的健壮性.比如你网站上针对最热门的文章或啥音乐电影之类的进行排名.由于这里排名不会像我们成绩排名会有并列第几名之说.所以出现了元素相等时也会有先后之分.如果添加进新的元素之后又要重新排名了.之前并列名次的最好是依然保持先后顺序才比较好.

## [递归算法的时间复杂度和空间复杂度](https://zhuanlan.zhihu.com/p/359006140)

## 排序简介

「排序算法 Sorting Algorithm」使得列表中的所有元素按照从小到大的顺序排列。

- 待排序的列表的 **元素类型** 可以是整数、浮点数、字符、或字符串；
- 排序算法可以根据需要设定 **判断规则** ，例如数字大小、字符 ASCII 码顺序、自定义规则；

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671975368421-b5edc6c7-9609-438c-b69d-29c95560d749.png)

Fig. 排序中的不同元素类型和判断规则

### 评价维度

排序算法主要可根据 **稳定性 、就地性 、自适应性 、比较类** 来分类。

### 稳定性

- 「稳定排序」在完成排序后，**不改变** 相等元素在数组中的相对顺序。
- 「非稳定排序」在完成排序后，相等元素在数组中的相对位置 **可能被改变**。

假设我们有一个存储学生信息当表格，第 1, 2 列分别是姓名和年龄。那么在以下示例中，「非稳定排序」会导致输入数据的有序性丢失。因此「稳定排序」是很好的特性，**在多级排序中是必须的**。

\# 输入数据是按照姓名排序好的 # (name, age)   ('A', 19)   ('B', 18)   ('C', 21)   ('D', 19)   ('E', 23)  # 假设使用非稳定排序算法按年龄排序列表， # 结果中 ('D', 19) 和 ('A', 19) 的相对位置改变， # 输入数据按姓名排序的性质丢失   ('B', 18)   ('D', 19)   ('A', 19)     ('C', 21)   ('E', 23) 

### 就地性

- 「原地排序」无需辅助数据，不使用额外空间；
- 「非原地排序」需要借助辅助数据，使用额外空间；

「原地排序」不使用额外空间，可以节约内存；并且一般情况下，由于数据操作减少，原地排序的运行效率也更高。

### 自适应性

- 「自适应排序」的时间复杂度受输入数据影响，即最佳 / 最差 / 平均时间复杂度不相等。
- 「非自适应排序」的时间复杂度恒定，与输入数据无关。

我们希望 **最差 = 平均** ，即不希望排序算法的运行效率在某些输入数据下发生劣化。

### 比较类

- 「比较类排序」基于元素之间的比较算子（小于、相等、大于）来决定元素的相对顺序。
- 「非比较类排序」不基于元素之间的比较算子来决定元素的相对顺序。

「比较类排序」的时间复杂度最优为 O(nlog⁡n) ；而「非比较类排序」可以达到 O(n) 的时间复杂度，但通用性较差。

### 理想排序算法

- **运行地快**，即时间复杂度低；
- **稳定排序**，即排序后相等元素的相对位置不变化；
- **原地排序**，即运行中不使用额外的辅助空间；
- **正向自适应性**，即算法的运行效率不会在某些输入数据下发生劣化；

然而，**没有排序算法同时具备以上所有特性**。排序算法的选型使用取决于具体的列表类型、列表长度、元素分布等因素。

## 交换排序

交换排序的**基本思想**是：两两比较待排序记录的关键字，一旦发现两个记录不满足次序要求时则进行交换，直到整个序列全部满足要求为止。

## 交换排序—[冒泡排序](https://www.topgoer.cn/docs/goalgorithm/goalgorithm-1cm6ato3141oq)

### 基本介绍

1. 冒泡排序(Bubble Sort)是一种最简单的交换排序方法，它通过两两比较相邻记录的关键字，如果发生逆序，则进行交换，从而使关键字小的记录如气泡一般逐渐往上 ＂漂浮＂（左移），或者使关键字大的记录如石块一样逐渐向下 ＂坠落”（右移）。
2. 每趟结束时，不仅能挤出一个最大值到最后面位置，还能同时部分理顺其他元素；           
3. 一旦下趟没有交换，还可提前结束排序

### 思想：

冒泡排序(Bubble Sorting）的**基本思想是**:

1. 通过对待排序序列从前向后（从下标较小的元素开始），依次比较相邻元素的值，
2. 若发现逆序则交换，使值较大的元素逐渐从前部移向后部，
3. 就象水底下的气泡一样逐渐向上冒。

注意：优化

**因为排序的过程中，各元素不断接近自己的位置，如果一趟比较下来没有进行过交换,就说明序列有序，因此要在排序过程中设置一个标志flag判断元素是否进行过交换。从而减少不必要的比较。**

### 算法步骤

- 从数组开头选择相邻两个元素进行比较，并进行交换
- 不停向后移动

1. 设待排序的记录存放在数组r[ 1 …n]中。首先将第一个记录的关键字和第二个记录的关键字进行比较，若为逆序（即 L.r[l].key>L.r[2].key), 则交换两个记录。然后比较第二个记录和第三个记录的关键字。依次类推，直至第n-1个记录和第n个记录的关键字进行过比较为止。上述过程称作第一趟起泡排序，其结果使得关键字最大的记录被安置到最后一个记录的位置上。
2. 然后进行第二趟起泡排序，对前n-1个记录进行同样操作，其结果是使关键字次大的记录被安置到第n — 1个记录的位置上。
3. 重复上述比较和交换过程，第 i 趟是从 L.r[1] 到 L.r[n-i+1] 依次比较相邻两个记录的关键字，并在 “逆序” 时交换相邻记录，其结果是这 n-i+1 个记录中关键字最大的记录被交换到第n-i+1 的位置上。直到在某一趟排序过程中没有进行过交换记录的操作，说明序列已全部达到排序要求，则完成排序。

### 思路分析图

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671802117667-9fb1946b-b0ba-43e0-afc6-348e7145fb1d.png)

待排序的记录总共有8个， 但算法在第六趟排序过程中没有进行过交换记录的操作，则完成排序。

### 代码实现

```go
package main

import "fmt"

func BubbleSort(list []int) {
	n := len(list)
	// 在一轮中有没有交换过
	didSwap := false

	// 进行 N-1 轮迭代
	for i := n - 1; i > 0; i-- {
		// 每次从第一位开始比较，比较到第 i 位就不比较了，因为前一轮该位已经有序了
		for j := 0; j < i; j++ {
			// 如果前面的数比后面的大，那么交换
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}

		// 如果在一轮中没有交换过，那么已经排好序了，直接返回
		if !didSwap {
			return
		}
	}
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	BubbleSort(list)
	fmt.Println(list)
}
```

### 复杂度

#### 时间复杂度

1. **最好：**情况（初始序列为正序）：只需进行一趟排序， 在排序过程中进行 **n-1** 次关键字间的比较，且**不移动记录，复杂度：O(n)**。
2. **最坏：**清况（初始序列为逆序）：需进行 n-1 趟排序，总的关键字比较次数 KCN和记录移动次数RMN(每次交换都要移动 3 次记录）分别为：

1. 1. ![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671803543221-3d70248e-0330-45f4-a546-df18a293fa9e.png)
   2. ![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671803550278-7fdc59dc-fff0-4947-ac0f-2138a6175fc1.png)
   3. 复杂度为**O(n^2**)

1. **平均：**情况下，冒泡排序关键字的**比较次数**和**记录移动**次数分别约为的 **n^2****/4**, 和 **n^2****/4**, 时间复杂度为**O(n^2**)。

#### 空间复杂度

冒泡排序只有在两个记录交换位置时需要一个辅助空间用做暂存记录，所以空间复杂度为**O(1)**。

### 稳定性

冒泡排序算法是稳定的，因为如果两个相邻元素相等，是不会交换的，保证了稳定性的要求。

### 算法特点

(1) 稳定排序。

(2) 可用于链式存储结构。

(3) 移动记录次数较多，算法平均时间性能比直接插入排序差。当初始记录无序，n较大时，此算法不宜采用。

### 适用场景

冒泡排序思路简单，代码也简单，特别适合小数据的排序。但是，由于算法复杂度较高，在数据量大的时候不适合使用。

## 交换排序—[快速排序](https://www.topgoer.cn/docs/goalgorithm/goalgorithm-1cm6b09u217l1#gcrdaw)

### 基本介绍

1. **课本描述：**快速排序 (Quick Sort) 是由冒泡排序改进而得的。 在 冒泡排序过程中， 只对相邻的两个记录进行比较， 因此每次交换两个相邻记录时只能消除一个逆序。 如果能通过两个（不相邻）记录的一次交换，消除多个逆序， 则会大大加快排序的速度。 快速排序方法中的一次交换可能消除多个逆序。
2. 快速排序属于**内部排序法**，也属于**交换类**的排序算法。
3. 其利用对问题的二分化，实现递归完成快速排序 ，在所有算法中二分化是最常用的方式，将问题尽量的分成两种情况加以分析， 最终以形成类似树的方式加以利用，因为在比较模型中的算法中，最快的排序时间 负载度为 **O(nlgn)**.
4. 快速排序是一种分治策略的排序算法，分治的理念就是依靠着递归。
5. 快速排序是一个知名度极高的排序算法，其对于大数据的优秀排序性能和相同复杂度算法中相对简单的实现使它注定得到比其他算法更多的宠爱。

### 思想

**基本思想**是:

1. 通过一趟排序将要排序的数据分割成独立的两部分，
2. 其中一部分的所有数据都比另外一部分的所有数据都要小，
3. 然后再按此方法对这两部分数据分别进行快速排序，
4. 整个排序过程可以递归进行，以此达到整个数据变成有序序列

### 算法步骤

1. 在待排序的 n个记录中任取一个记录（通常取第一个记录）作为枢轴（或支点），设其关键字为pivotkey。
2. 经过一趟排序后，把所有关键字小于pivotkey 的记录交换到前面，把所有关键字大于pivotkey的记录交换到后面，结果将待排序记录分成两个子表，最后将枢轴放置在分界处的位置。
3. 然后，分别对左、右子表重复上述过程，直至每一子表只有一个记录时，排序完成。

其中，**一趟**快速排序的**具体步骤**如下。

1. 选择待排序表中的第一个记录作为枢轴，将枢轴记录暂存在 r[0]的位置上。附设两个指针 low和 high, 初始时分别指向表的下界和上界（第一趟时， low = 1; high= L.length;)。
2. 从表的最右侧位置依次向左搜索 ，找到第一个关键字小于枢轴关键字 pivotkey 的记录，将其移到 low 处。

1. 1. 具体操作是：当 low<high 时，若 high所指记录的关键字大千等于 pivotkey, 则向左移动指针 high (执行操作--high); 
   2. 否则将 high所指记录移到 low所指记录。

1. 然后再从表的最左侧位置，依次向右搜索找到第一个关键字大于 pivotkey 的记录和枢轴记录交换。

1. 1. 具体操作是：当 low<high 时，若 low所指记录的关键字小于等于 pivotkey, 则向右移动指针 low (执行操作++low); 
   2. 否则将 low所指记录与枢轴记录交换。

1. 复步骤2和3， 直至 low 与 high 相等为止。此时 low 或 high 的位置即为枢轴在此趟排序中的最终位置， 原表被分成两个子表。

在 上述过程中， 记录的交换都是与枢轴之间发生， 每次交换都要移动 3次记录， 可以先将枢轴记录暂存在 r[0]的位置上， 排序过程中只移动要与枢轴交换的记录， 即只做 r[low]或 r[high]的单向移动， 直至一趟排序结束后再将枢轴记录移至正确位置上。

### 思路分析图

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671704854368-84c6f27a-d95b-4268-acc6-bb1dc5fc40a1.png)

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671704797953-9c01ce0d-9eac-4eac-aa62-c026259317c7.png)

### 代码实现

```go
package main

import (
	"fmt"
)

/**
 * @Author huChao
 * @Description 交换排序-快速排序
 * @Date 20:56 2022/12/21
 * @Param 1. left 表示 数组左边的下标
 * @Param 2. right 表示数组右边的下标
 * @Param 3. array 表示要排序的数组
 * @return []int
 **/

func QuickSort(left int, right int, array []int) []int {
	l := left
	r := right
	// pivot 是中轴， 支点
	pivot := array[(left+right)/2]
	//for 循环的目标是将比 pivot 小的数放到 左边
	//  比 pivot 大的数放到 右边
	for l < r {
		//从  pivot 的左边找到大于等于pivot的值
		for array[l] < pivot {
			l++
		}
		//从  pivot 的右边边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		// 1 >= r 表明本次分解任务完成, break
		if l >= r {
			break
		}
		//交换
		array[l], array[r] = array[r], array[l]
		//优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	// 如果  1== r, 再移动下
	if l == r {
		l++
		r--
	}
	// 向左递归
	if left < r {
		QuickSort(left, r, array)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, array)
	}
	return array
}

func main() {
	arr := []int{-9, 78, 0, 23, -567, 70}
	quickSort := QuickSort(0, len(arr)-1, arr)
	fmt.Println(quickSort) // [-567 -23 -9 0 23 70]
}
```

### 复杂度

#### 时间复杂度

- **最好：**在最好情况下，每一轮都能平均切分，这样遍历元素只要 n/2 次就可以把数列分成两部分，每一轮的时间复杂度都是：O(n)。因为问题规模每次被折半，折半的数列继续递归进行切分，也就是总的时间复杂度计算公式为： T(n) = 2*T(n/2) + O(n)。按照主定理公式计算，我们可以知道时间复杂度为：**O(nlogn)**，当然我们可以来具体计算一下：

```go
我们来分析最好情况，每次切分遍历元素的次数为 n/2

T(n) = 2*T(n/2) + n/2
T(n/2) = 2*T(n/4) + n/4
T(n/4) = 2*T(n/8) + n/8
T(n/8) = 2*T(n/16) + n/16
...
T(4) = 2*T(2) + 4
T(2) = 2*T(1) + 2
T(1) = 1

进行合并也就是：

T(n) = 2*T(n/2) + n/2
     = 2^2*T(n/4)+ n/2 + n/2
     = 2^3*T(n/8) + n/2 + n/2 + n/2
     = 2^4*T(n/16) + n/2 + n/2 + n/2 + n/2
     = ...
     = 2^logn*T(1) + logn * n/2
     = 2^logn + 1/2*nlogn
     = n + 1/2*nlogn

因为当问题规模 n 趋于无穷大时 nlogn 比 n 大，所以 T(n) = O(nlogn)。

最好时间复杂度为：O(nlogn)。
```

- **最差：**最差的情况下，每次都不能平均地切分，每次切分都因为基准数是最大的或者最小的，不能分成两个数列，这样时间复杂度变为了 T(n) = T(n-1) + O(n)，按照主定理计算可以知道时间复杂度为：**O(n^2**)，我们可以来实际计算一下：

```go
我们来分析最差情况，每次切分遍历元素的次数为 n

T(n) = T(n-1) + n
     = T(n-2) + n-1 + n
     = T(n-3) + n-2 + n-1 + n
     = ...
     = T(1) + 2 +3 + ... + n-2 + n-1 + n
     = O(n^2)

最差时间复杂度为：O(n^2)。
```

- **平均：**根据熵的概念，数量越大，随机性越高，越自发无序，所以待排序数据规模非常大时，出现最差情况的情形较少。在综合情况下，快速排序的平均时间复杂度为：**O(nlogn)**。实验结果表明：就平均计算时间而言，快速排序是我们所讨论的所有内排序方法中最好的一个。

#### 空间复杂度

1. 快速排序是递归的，需要有一个栈存放每层递归调用时参数（新的low和high），执行时需要有一个栈来存放相应的数据。
2. 最大递归调用次数与递归树的深度一致，所以**最好**情况下的空间复杂度为O(log2n)，
3. **最坏**情况下为，若每次只完成了一个元素，那么空间复杂度为 O(n)。
4. 所以我们一般认为快速排序的空间复杂度为 O(log2n)。

### 稳定性

快速排序是不稳定的，因为切分过程中进行了交换，相同值的元素可能发生位置变化。

### 算法特点

1. 记录非顺次的移动导致排序方法是不稳定的。
2. 排序过程中需要定位表的下界和上界，所以适合用于顺序结构，很难用于链式结构。
3. 当n较大时，在平均情况下快速排序是所有内部排序方法中速度最快的一种，所以其适合初始记录无序、 n较大时的情况。

### 适用场景

快速排序在大多数情况下都是适用的，尤其在数据量大的时候性能优越性更加明显。但是在必要的时候，需要考虑下优化以提高其在最坏情况下的性能。

### 补充：内置库使用快速排序的原因

首先堆排序，归并排序最好最坏时间复杂度都是：O(nlogn)，而快速排序最坏的时间复杂度是：O(n^2)，但是很多编程语言内置的排序算法使用的仍然是快速排序，这是为什么？

1. 这个问题有偏颇，选择排序算法要看具体的场景，Linux 内核用的排序算法就是堆排序，而 Java 对于数量比较多的复杂对象排序，内置排序使用的是归并排序，只是一般情况下，快速排序更快。
2. 归并排序有两个稳定，第一个稳定是排序前后相同的元素位置不变，第二个稳定是，每次都是很平均地进行排序，读取数据也是顺序读取，能够利用存储器缓存的特征，比如从磁盘读取数据进行排序。因为排序过程需要占用额外的辅助数组空间，所以这部分有代价损耗，但是原地手摇的归并排序克服了这个缺陷。
3. 复杂度中，大 O 有一个常数项被省略了，堆排序每次取最大的值之后，都需要进行节点翻转，重新恢复堆的特征，做了大量无用功，常数项比快速排序大，大部分情况下比快速排序慢很多。但是堆排序时间较稳定，不会出现快排最坏 O(n^2) 的情况，且省空间，不需要额外的存储空间和栈空间。
4. 当待排序数量大于16000个元素时，使用自底向上的堆排序比快速排序还快，可见此：https://core.ac.uk/download/pdf/82350265.pdf。
5. 快速排序最坏情况下复杂度高，主要在于切分不像归并排序一样平均，而是很依赖基准数的现在，我们通过改进，比如随机数，三切分等，这种最坏情况的概率极大的降低。大多数情况下，它并不会那么地坏，大多数快才是真的块。
6. 归并排序和快速排序都是分治法，排序的数据都是相邻的，而堆排序比较的数可能跨越很大的范围，导致局部性命中率降低，不能利用现代存储器缓存的特征，加载数据过程会损失性能。

对稳定性有要求的，要求排序前后相同元素位置不变，可以使用归并排序，Java 中的复杂对象类型，要求排序前后位置不能发生变化，所以小规模数据下使用了直接插入排序，大规模数据下使用了归并排序。

对栈，存储空间有要求的可以使用堆排序，比如 Linux 内核栈小，快速排序占用程序栈太大了，使用快速排序可能栈溢出，所以使用了堆排序。

## 选择排序

## 选择排序—简单[选择排序](https://www.topgoer.cn/docs/goalgorithm/goalgorithm-1cm6att75q48j)

### 基本介绍

选择式排序也属于内部排序法，是从欲排序的数据中，按指定的规则选出某一元素，经过和其他元素重整，再依原则交换位置后达到排序的目的。

**其实选择排序是非常简单的，和冒泡排序有异曲同工之妙。就是把元素分成两部分，一部分是有序的，另外一部分是无序的；每次循环从无序的元素中选取一个元素放到有序的元素中，依次循环到最后把所有元素都放到了有序那一部分中（也就是无序部分，元素为零）；**

### 思想：

选择排序（select sorting）也是一种简单的排序方法。它的**基本思想**是:

1. 第一次从R[0]~R[n-1]中选取最小值，与R[0]交换，
2. 第二次从R[1]~R[n-1]中选取最小值，与R[1]交换，
3. 第三次从R[2]~R[n-1]中选取最小值，与R[2]交换，…，
4. 第i次从R[i-1]~R[n-1]中选取最小值，与R[-1]交换，…，
5. 第n-1次从R[n-2]~R[n-1]中选取最小值，与R[n-2]交换，
6. 总共通过n-1次，得到一个按排序码从小到大排列的有序序列。

### 算法步骤

**描述一：**

1. 在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
2. 从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
3. 重复第二步，直到所有元素均排序完毕。

**描述二：**

1、外循环：循环每个位置（其实就是选择了这个位置，然后用内循环去选择一个合适的数，放到这个位置）；

2、内循环：在无序元素中选择一个合适的数；

3、把第二步选中的数据放到第一步选中的位置上就可以了；

### 选择排序思路分析图

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671459717188-c9be4e71-801e-4e76-af3a-f2a9d5d1b42b.png)

### 代码实现

```go
package main

import (
	"fmt"
)

//SelectSort 选择排序
// 最好、最坏、平均时间复杂度均为：O(n^2),
// 空间复杂度:O(1)
// 稳定性：不稳定 如:5 5 2
func SelectSort(arr []int) []int {
	//1. 先完成将第一个最大值和 arr[0] => 先易后难
	//1 假设  arr[0] 最大值
	for j := 0; j < len(arr)-1; j++ {

		max := arr[j]
		maxIndex := j
		//2. 遍历后面 1---[len(arr) -1] 比较
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] { //找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		fmt.Printf("第%d次 %v\n  ", j+1, arr)
	}
	return arr
}

func main() {
	//定义一个数组
	arr := []int{10, 34, 19, 100, 80, 789}
	selectSort := SelectSort(arr)
	fmt.Println(selectSort)
}

/**
第1次 [789 34 19 100 80 10]
  第2次 [789 100 19 34 80 10]
  第3次 [789 100 80 34 19 10]
  第4次 [789 100 80 34 19 10]
  第5次 [789 100 80 34 19 10]
  [789 100 80 34 19 10] 
*/
```

### 复杂度

#### 时间复杂度

可以很直观的看出选择排序的时间复杂度：就是两个循环消耗的时间；

- 第一次内循环比较N - 1次，然后是N-2次，N-3次，……，最后一次内循环比较1次。共比较的次数是 (N - 1) + (N - 2) + ... + 1，求等差数列和，得 (N - 1 + 1)* N / 2 = N^2 / 2。舍去最高项系数，其时间复杂度为 O(N^2)。
- 所以**最优**的时间复杂度  和**最差**的时间复杂度   和**平均**时间复杂度  都为 ：**O(n^2)**

#### 空间复杂度，

- 最优的情况下（已经有顺序）复杂度为：O(0) ；
- 最差的情况下（全部元素都要重新排序）复杂度为：O(n );
- 平均的时间复杂度：O(1)

### 稳定性

选择排序是一个不稳定的排序算法，比如数组：[5 6 5 1]，第一轮迭代时最小的数是 1，那么与第一个元素 5 交换位置，这样数字 1 就和数字 5 交换了位置，导致两个相同的数字 5 排序后位置变了。

### 适用场景

选择排序实现也比较简单，并且由于在各种情况下复杂度波动小，因此一般是优于冒泡排序的。在所有的完全交换排序中，选择排序也是比较不错的一种算法。但是，由于固有的O(n2)复杂度，选择排序在海量数据面前显得力不从心。因此，它适用于简单数据排序。

------

## 选择排序—树形选择排序

## 选择排序—[堆排序](https://www.topgoer.com/Go高级/堆排序算法.html)

### 原理

是对直接选择排序的改进，不稳定，时间复杂度 O(nlogn)，空间复杂度 O(1)。

将待排序记录看作完全⼆叉树，可以建⽴⼤根堆或⼩根堆，⼤根堆中每个节点的值都不⼩于它的⼦节点 值，⼩根堆中每个节点的值都不⼤于它的⼦节点值。

以⼤根堆为例，在建堆时⾸先将最后⼀个节点作为当前节点，如果当前节点存在⽗节点且值⼤于⽗节点，就将当前节点和⽗节点交换。在移除时⾸先暂存根节点的值，然后⽤最后⼀个节点代替根节点并作 为当前节点，如果当前节点存在⼦节点且值⼩于⼦节点，就将其与值较⼤的⼦节点进⾏交换，调整完堆 后返回暂存的值。



### 实现

算法描述：首先建一个堆，然后调整堆，调整过程是将节点和子节点进行比较，将 其中最大的值变为父节点，递归调整调整次数lgn,最后将根节点和尾节点交换再n次 调整**O(nlgn)**.

#### 算法步骤

- 创建最大堆或者最小堆（我是最小堆）
- 调整堆
- 交换首尾节点(为了维持一个完全二叉树才要进行收尾交换)

```go
package sort

import "fmt"

//堆排序
func main() {
    arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
    fmt.Println(HeapSort(arr))
}
func HeapSortMax(arr []int, length int) []int {
    // length := len(arr)
    if length <= 1 {
        return arr
    }
    depth := length/2 - 1 //二叉树深度
    for i := depth; i >= 0; i-- {
        topmax := i //假定最大的位置就在i的位置
        leftchild := 2*i + 1
        rightchild := 2*i + 2
        if leftchild <= length-1 && arr[leftchild] > arr[topmax] { //防止越过界限
            topmax = leftchild
        }
        if rightchild <= length-1 && arr[rightchild] > arr[topmax] { //防止越过界限
            topmax = rightchild
        }
        if topmax != i {
            arr[i], arr[topmax] = arr[topmax], arr[i]
        }
    }
    return arr
}
func HeapSort(arr []int) []int {
    length := len(arr)
    for i := 0; i < length; i++ {
        lastlen := length - i
        HeapSortMax(arr, lastlen)
        if i < length {
            arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
        }
    }
    return arr
}
```

## 插入排序

## 插入排序—直接[插入排序](https://www.topgoer.cn/docs/goalgorithm/goalgorithm-1cm6avesshjb1)



### 基本介绍

插入式排序属于内部排序法，是对于欲排序的元素以插入的方式找寻该元素的适当位置，以达到排序的目的。

### 思想

插入排序（Insertion Sorting)的**基本思想**是:

1. 把n个待排序的元素看成为一个有序表和一个无序表
2. 开始时有序表中只包含一个元素，无序表中包含有n-1个元素
3. 排序过程中每次从无序表中取出第一个元素，把它的排序码依次与有序表元素的排序码进行比较，将它插入到有序表中的适当位置，使之成为新的有序表。

### 算法步骤

1. 把待排序的数组分成已排序和未排序两部分，初始的时候把第一个元素认为是已排好序的。
2. 从第二个元素开始，在已排好序的子数组中寻找到该元素合适的位置并插入该位置。
3. 重复上述过程直到最后一个元素被插入有序子数组中。

### 思路分析图

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1671528532861-7cf60c5c-e9f0-4eea-94e3-1bdd20553911.png)

### 代码实现

```go
package main

import (
	"fmt"
)

// InsertSort 插入排序
// 平均时间复杂度o(n^2)
// 最好时间复杂度o(n)
// 空间复杂度o(1)
// 稳定
func InsertSort(arr []int) []int {

	//完成第一次，给第二个元素找到合适的位置并插入
	for i := 1; i < len(arr); i++ {

		insertVal := arr[i]
		insertIndex := i - 1 // 下标

		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次插入后 %v\n", i, arr)
	}
	return arr
}
func main() {
	arr := []int{23, 0, 12, 56, 34, -1, 55}
	insertSort := InsertSort(arr)
	fmt.Println(insertSort)
}

/**
第1次插入后 [23 0 12 56 34 -1 55]
第2次插入后 [23 12 0 56 34 -1 55]
第3次插入后 [56 23 12 0 34 -1 55]
第4次插入后 [56 34 23 12 0 -1 55]
第5次插入后 [56 34 23 12 0 -1 55]
第6次插入后 [56 55 34 23 12 0 -1]
[56 55 34 23 12 0 -1]
*/
```

### 复杂度

#### 时间复杂度

1. **平均：**时间复杂度 O(n²)
2. **最坏：**插入排序的时间复杂度分析。在最坏情况下，数组完全逆序，插入第2个元素时要考察前1个元素，插入第3个元素时，要考虑前2个元素，……，插入第N个元素，要考虑前 N - 1 个元素。因此，最坏情况下的比较次数是 1 + 2 + 3 + ... + (N - 1)，等差数列求和，结果为 N^2 / 2，所以最坏情况下的复杂度为 O(N^2)。
3. **最好**：最坏情况下，数组已经是有序的，每插入一个元素，只需要考查前一个元素，因此最好情况下，插入排序的时间复杂度为O(N)

#### 空间复杂度

只使用了i,insertIndex,insertVal这两个辅助元素，与问题规模无关，空间复杂度为O(1)

### 稳定性

因为是从右到左，将一个个未排序的数，插入到左边已排好序的队列中，所以插入排序，相同的数在排序后顺序不会变化，这个排序算法是稳定的。

### 适用场景

插入排序由于O( n2 )的复杂度，在数组较大的时候不适用。但是，在数据比较少的时候，是一个不错的选择，一般做为快速排序的扩充。

1. 数组规模 n 较小的大多数情况下，我们可以使用插入排序，它比冒泡排序，选择排序都快，甚至比任何的排序算法都快。
2. 数列中的有序性越高，插入排序的性能越高，因为待排序数组有序性越高，插入排序比较的次数越少。
3. 大家都很少使用冒泡、直接选择，直接插入排序算法，因为在有大量元素的无序数列下，这些算法的效率都很低。

## 插入排序—折半排序

## 插入排序—希尔排序

⼜称缩⼩增量排序，是对直接插⼊排序的改进，不稳定，平均时间复杂度 O(n^1.3^)，最差时间复杂度O(n²)，最好时间复杂度 O(n)，空间复杂度 O(1)。

把记录按下标的⼀定增量分组，对每组进⾏直接插⼊排序，每次排序后减⼩增量，当增量减⾄ 1 时排序完毕。

## [归并排序](https://hunterhug.gitlab.io/goa.c/#/algorithm/sort/merge_sort)—内部排序

### 基本介绍

1. 归并排序是一种分治策略的排序算法。
2. 它是一种比较特殊的排序算法，通过递归地先使每个子序列有序，再将两个有序的序列进行合并成一个有序的序列。
3. **归并排序(Merging Sort)**就是将两个或两个以上的有序表合并成一个有序表的过程。将两个有序表合并成一个有序表的过程称为**2-路归并**，2-路归并 最为简单和常用。

### 思想

「归并排序 Merge Sort」是算法中“分治思想”的典型体现，其有「划分」和「合并」两个阶段：

1. **划分阶段：** 通过递归不断 **将数组从中点位置划分开**，将长数组的排序问题转化为短数组的排序问题；
2. **合并阶段：** 划分到子数组长度为 1 时，开始向上合并，不断将 **左、右两个短排序数组** 合并为 **一个长排序数组**，直至合并至原数组时完成排序；

### 算法步骤

**「递归划分」** 从顶至底递归地 **将数组从中点切为两个子数组** ，直至长度为 1 ；

1. 计算数组中点 mid ，递归划分左子数组（区间 [left, mid] ）和右子数组（区间 [mid + 1, right] ）；
2. 递归执行 1. 步骤，直至子数组区间长度为 1 时，终止递归划分；

**「回溯合并」** 从底至顶地将左子数组和右子数组合并为一个 **有序数组** ；

需要注意，由于从长度为 1 的子数组开始合并，所以 **每个子数组都是有序的** 。因此，合并任务本质是要 **将两个有序子数组合并为一个有序数组** 。

观察发现，归并排序的递归顺序就是二叉树的「后序遍历」。

- **后序遍历：** 先递归左子树、再递归右子树、最后处理根结点。
- **归并排序：** 先递归左子树、再递归右子树、最后处理合并。

### 思路分析图


![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1672062301725-1c8e0d4c-49c1-4438-b560-30ce7a95b2b8.png)

### 代码实现

```go
package main

import "fmt"

// 合并左子数组和右子数组
// 左子数组区间 [left, mid]
// 右子数组区间 [mid + 1, right]
func merge(nums []int, left, mid, right int) {
	// 初始化辅助数组 借助 copy模块
	tmp := make([]int, right-left+1)
	for i := left; i <= right; i++ {
		tmp[i-left] = nums[i]
	}
	// 左子数组的起始索引和结束索引
	leftStart, leftEnd := left-left, mid-left
	// 右子数组的起始索引和结束索引
	rightStart, rightEnd := mid+1-left, right-left
	// i, j 分别指向左子数组、右子数组的首元素
	i, j := leftStart, rightStart
	// 通过覆盖原数组 nums 来合并左子数组和右子数组
	for k := left; k <= right; k++ {
		// 若“左子数组已全部合并完”，则选取右子数组元素，并且 j++
		if i > leftEnd {
			nums[k] = tmp[j]
			j++
			// 否则，若“右子数组已全部合并完”或“左子数组元素 < 右子数组元素”，则选取左子数组元素，并且 i++
		} else if j > rightEnd || tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
			// 否则，若“左子数组元素 > 右子数组元素”，则选取右子数组元素，并且 j++
		} else {
			nums[k] = tmp[j]
			j++
		}
	}
}

func mergeSort(nums []int, left, right int) {
	// 终止条件
	if left >= right {
		return
	}
	// 划分阶段
	mid := (left + right) / 2
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	// 合并阶段
	merge(nums, left, mid, right)
}

func main() {
	nums := []int{7, 3, 2, 6, 0, 1, 5, 4}
	mergeSort(nums, 0, len(nums)-1)
	fmt.Println("归并排序完成后 nums = ", nums)
}
```

### 下面重点解释一下合并方法 merge() 的流程：

1. 初始化一个辅助数组 tmp 暂存待合并区间 [left, right] 内的元素，后续通过覆盖原数组 nums 的元素来实现合并；
2. 初始化指针 i , j , k 分别指向左子数组、右子数组、原数组的首元素；
3. 循环判断 tmp[i] 和 tmp[j] 的大小，将较小的先覆盖至 nums[k] ，指针 i , j 根据判断结果交替前进（指针 k 也前进），直至两个子数组都遍历完，即可完成合并。

合并方法 merge() 代码中的主要难点：

- nums 的待合并区间为 [left, right] ，而因为 tmp 只复制了 nums 该区间元素，所以 tmp 对应区间为 [0, right - left] ，**需要特别注意代码中各个变量的含义**。
- 判断 tmp[i] 和 tmp[j] 的大小的操作中，还 **需考虑当子数组遍历完成后的索引越界问题**，即 i > leftEnd 和 j > rightEnd 的情况，索引越界的优先级是最高的，例如如果左子数组已经被合并完了，那么不用继续判断，直接合并右子数组元素即可。

### 复杂度

#### 时间复杂度

每次都是一分为二，特别均匀，所以**最差和最坏**时间复杂度**都一样**。归并操作的时间复杂度为：O(n)，因此总的时间复杂度为：T(n)=2T(n/2)+O(n)，根据主定理公式可以知道时间复杂度为：**O(nlogn)**。我们可以自己计算一下：

```go
归并排序，每次归并操作比较的次数为两个有序数组的长度： n/2

T(n) = 2*T(n/2) + n/2
T(n/2) = 2*T(n/4) + n/4
T(n/4) = 2*T(n/8) + n/8
T(n/8) = 2*T(n/16) + n/16
...
T(4) = 2*T(2) + 4
T(2) = 2*T(1) + 2
T(1) = 1

进行合并也就是：

T(n) = 2*T(n/2) + n/2
     = 2^2*T(n/4)+ n/2 + n/2
     = 2^3*T(n/8) + n/2 + n/2 + n/2
     = 2^4*T(n/16) + n/2 + n/2 + n/2 + n/2
     = ...
     = 2^logn*T(1) + logn * n/2
     = 2^logn + 1/2*nlogn
     = n + 1/2*nlogn

因为当问题规模 n 趋于无穷大时 nlogn 比 n 大，所以 T(n) = O(nlogn)。

因此时间复杂度为：O(nlogn)。
```

#### 空间复杂度 **O(n)** **：** 

1. 用顺序表实现归并排序时， 需要和待排序记录个数相等的辅助存储空间， 所以空间复杂度为O(n)。
2. 递归深度为 log⁡n ，使用 O(log⁡n) 大小的栈帧空间。

### 稳定性

在合并时可保证相等元素的相对位置不变。

### 算法特点

- **非原地排序：** 辅助数组需要使用 O(n) 额外空间。
- **非自适应排序：** 对于任意输入数据，归并排序的时间复杂度皆相同。

### 适用场景

1. 归并排序在数据量比较大的时候也有较为出色的表现（效率上），但是，其空间复杂度O(n)使得在数据量特别大的时候（例如，1千万数据）几乎不可接受。而且，考虑到有的机器内存本身就比较小，因此，采用归并排序一定要注意。
2. 归并排序有一个很特别的优势，用于排序链表时有很好的性能表现，**空间复杂度可被优化至 O(1)** ，这是因为：

- 由于链表可仅通过改变指针来实现结点增删，因此“将两个短有序链表合并为一个长有序链表”无需使用额外空间，即回溯合并阶段不用像排序数组一样建立辅助数组 tmp ；
- 通过使用「迭代」代替「递归划分」，可省去递归使用的栈帧空间；

## [基数排序算法](https://www.topgoer.com/Go高级/基数排序算法.html)



算法描述：基数排序类似计数排序，需要额外的空间来记录对应的基数内的数据 额外的空间是有序的，最终时间复杂度**O(nlogrm)**,r是基数，r^m=n.当给定 特定的范围，计数排序又可以叫桶排序，当以10进制为基数时就是简单的桶排序

#### 算法步骤

- 从个位开始排序，从低到高进行递推
- 比较过程中如果遇到高位相同时，顺序不变

#### 算法分两类

1. 低位排序LSD
2. 高位排序MSD

```go
package sort

import "fmt"

func main() {
    var arr [3][]int
    myarr := []int{1, 2, 3, 1, 1, 2, 2, 2, 2, 2, 3}
    for i := 0; i < len(myarr); i++ {
        arr[myarr[i]-1] = append(arr[myarr[i]-1], myarr[i])
    }
    fmt.Println(arr)
}
```

## Q16：[拓扑排序](https://www.topgoer.com/Go高级/拓扑排序.html)

####  定义

对于一些有前后依赖关系的排序算法，是利用有向无环图进行实现，通过局部依赖关系确定全局顺序的算法

#### 应用场景

- 编译有序依赖的文件

### 1.1.1. 两种拓扑算法

#### Kahn算法

- 算法逻辑
- 利用贪心算法，如果两个顶点，顶点b依赖于顶点a,就将a指向b,当一个顶点的入度为零，将这个顶点就是最优排序点， 并且将顶点从图中移除，将可达顶点的入度减一。

#### DFS算法

1.使用深度算法，产生逆向邻接表先输出其他依赖，最后输出自己。

```go
package main

import (
    "fmt"
)

//有向图
type graph struct {
    vertex int           //顶点
    list   map[int][]int //连接表边
}

//添加边
func (g *graph) addVertex(t int, s int) {
    g.list[t] = push(g.list[t], s)
}

func main() {
    g := NewGraph(8)
    g.addVertex(2, 1)
    g.addVertex(3, 1)
    g.addVertex(7, 1)
    g.addVertex(4, 2)
    g.addVertex(5, 2)
    g.addVertex(8, 7)
    g.DfsSort()
}

//创建图
func NewGraph(v int) *graph {
    g := new(graph)
    g.vertex = v
    g.list = map[int][]int{}
    i := 0
    for i < v {
        g.list[i] = make([]int, 0)
        i++
    }
    return g
}

//取出切片第一个
func pop(list []int) (int, []int) {
    if len(list) > 0 {
        a := list[0]
        b := list[1:]
        return a, b
    } else {
        return -1, list
    }
}

//推入切片
func push(list []int, value int) []int {
    result := append(list, value)
    return result
}

//添加边
func (g *graph) KhanSort() {
    var inDegree = make(map[int]int)
    var queue []int
    for i := 1; i <= g.vertex; i++ {
        for _, m := range g.list[i] {
            inDegree[m]++
        }
    }
    for i := 1; i <= g.vertex; i++ {

        if inDegree[i] == 0 {
            queue = push(queue, i)
        }
    }
    for len(queue) > 0 {
        var now int
        now, queue = pop(queue)
        fmt.Println("->", now)
        for _, k := range g.list[now] {
            inDegree[k]--
            if inDegree[k] == 0 {
                queue = push(queue, k)
            }
        }
    }
}

func (g *graph) DfsSort() {
    inverseList := make(map[int][]int)
    //初始化逆向邻接表
    for i := 1; i <= g.vertex; i++ {
        for _, k := range g.list[i] {
            inverseList[k] = append(inverseList[k], i)
        }
    }
    visited := make([]bool, g.vertex+1)
    visited[0] = true
    for i := 1; i <= g.vertex; i++ {
        if visited[i] == false {
            visited[i] = true
            dfs(i, inverseList, visited)
        }
    }

}

func dfs(vertex int, inverseList map[int][]int, visited []bool) {
    for _, w := range inverseList[vertex] {
        if visited[w] == true {
            continue
        } else {
            visited[w] = true
            dfs(w, inverseList, visited)
        }
    }
    fmt.Println("->", vertex)
}
```

## Q16：循环和递归，你说下有什么不同的点？



递归算法：

优点：代码少、简介。

缺点：它的运行需要较多次数的函数调用，如果调用层数比较深，需要增加额外的堆栈处理，比如参数传递需要压栈等操作，会对执行效率有一定影响。但是，对于某些问题，如果不使用递归，那将是极端难看的代码。

循环算法：

优点：速度快，结构简单。

缺点：并不能解决所有的问题。有的问题适合使用递归而不是循环。如果使用循环并不困难的话，最好使用循环。

## Q17：排序算法怎么选择？

数据量规模较⼩，考虑直接插⼊或直接选择。当元素分布有序时直接插⼊将⼤⼤减少⽐较和移动记录的次数，如果不要求稳定性，可以使⽤直接选择，效率略⾼于直接插⼊。

数据量规模中等，选择希尔排序。

数据量规模较⼤，考虑堆排序（元素分布接近正序或逆序）、快速排序（元素分布随机）和归并排序稳定性）。⼀般不使⽤冒泡。

数据分布比较均匀，桶排序。



# 查找

## [二分查找](https://www.topgoer.com/Go高级/二分查找方法.html) 

算法描述：在一组有序数组中，将数组一分为二，将要查询的元素和分割点进行比较，分为三种情况

- 相等直接返回
- 元素大于分割点，在分割点右侧继续查找
- 元素小于分割点，在分割点左侧继续查找

时间复杂： **O(lgn)**.

#### 要求

- 必须是有序的数组，并能支持随机访问

#### 变形

- 查找第一个值等于给定的

- - 在相等的时候做处理，向前查

- 查找最后一个值等于给定的值

- - 在相等的时候做处理，向后查

- 查找第一个大于等于给定的值

- - 判断边界减1

- 查找最后一个小于等于给定的值

- - 判断边界加1

#### 实际应用

- 用户ip区间段查询
- 用于相似度查询



```go
package sort

import "fmt"

func bin_search(arr []int, finddata int) int {
    low := 0
    high := len(arr) - 1
    for low <= high {
        mid := (low + high) / 2
        fmt.Println(mid)
        if arr[mid] > finddata {
            high = mid - 1
        } else if arr[mid] < finddata {
            low = mid + 1
        } else {
            return mid
        }
    }
    return -1
}

func main() {
    arr := make([]int, 1024*1024, 1024*1024)
    for i := 0; i < 1024*1024; i++ {
        arr[i] = i + 1
    }
    id := bin_search(arr, 1024)
    if id != -1 {
        fmt.Println(id, arr[id])
    } else {
        fmt.Println("没有找到数据")
    }
}
```