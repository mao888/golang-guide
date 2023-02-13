package main

import "fmt"

// 1） 如果没有指定max：max的值为截取对象（数组、切片）的容量
//
// 2） 如果指定max：max 值不能超过原对象（数组、切片）的容量。
//
// 3）利用数组创建切片，切片操作的是同一个底层数组。

// 参考文章：https://blog.csdn.net/weixin_42117918/article/details/81913036

func main() {

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("原数组：", arr)

	fmt.Println("对数组进行截取：")
	//如果指定max，max的值最大不能超过截取对象（数组、切片）的容量
	s1 := arr[2:5:9] //max:9  low：2  high;5  len:5-2(len=high-low)  cap:9-2(cap=max-low)
	fmt.Printf("数组截取之后的类型为：%T,    数据是：%v;长度：%d;容量：%d\n", s1, s1, len(s1), cap(s1))
	// 数组截取之后的类型为：[]int,    数据是：[3 4 5];长度：3;容量：7

	//如果没有指定max，max的值为截取对象（切片、数组）的容量
	s2 := s1[1:7] //max:7  low：1  high;7  len:7-1(len=high-low)  cap:7-1(cap=max-low)
	fmt.Println("对切片进行截取：")
	fmt.Printf("对切片进行截取之后的数据是：%v,长度:%d； 容量%d\n", s2, len(s2), cap(s2))
	// 对切片进行截取之后的数据是：[4 5 6 7 8 9],长度:6； 容量6

	//利用数组创建切片，切片操作的是同一个底层数组
	s1[0] = 8888
	s2[0] = 6666
	fmt.Println("操作之后的数组为：", arr)
	// 操作之后的数组为： [1 2 8888 6666 5 6 7 8 9 10]
	/*
		切片对数组的截取  最终都是切片操作的底层数组（通过指针操作原数组）
	*/

}
