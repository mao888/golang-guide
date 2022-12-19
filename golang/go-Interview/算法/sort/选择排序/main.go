package main
import (
	"fmt"
	"math/rand"
	"time"
)

//编写函数selectSort 完成排序

func SelectSort(arr *[80000]int) {

	//标准的访问方式
	//(*arr)[1] = 600 等价于 arr[1] = 900
	//arr[1] = 900
	//1. 先完成将第一个最大值和 arr[0] => 先易后难

	//1 假设  arr[0] 最大值

	for j := 0; j < len(arr) - 1; j++ {

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

		//fmt.Printf("第%d次 %v\n  ", j+1 ,*arr)
	}
	

	/*
	max = arr[1]
	maxIndex = 1
	//2. 遍历后面 2---[len(arr) -1] 比较
	for i := 1 + 1; i < len(arr); i++ {
		if max < arr[i] { //找到真正的最大值
			max = arr[i]
			maxIndex = i
		}
	}
	//交换
	if maxIndex != 1 {
		arr[1], arr[maxIndex] = arr[maxIndex], arr[1]
	}

	fmt.Println("第2次 ", *arr)

	

	max = arr[2]
	maxIndex = 2
	//2. 遍历后面 3---[len(arr) -1] 比较
	for i := 2 + 1; i < len(arr); i++ {
		if max < arr[i] { //找到真正的最大值
			max = arr[i]
			maxIndex = i
		}
	}
	//交换
	if maxIndex != 2 {
		arr[2], arr[maxIndex] = arr[maxIndex], arr[2]
	}

	fmt.Println("第3次 ", *arr)

	max = arr[3]
	maxIndex = 3
	//2. 遍历后面 4---[len(arr) -1] 比较
	for i := 3 + 1; i < len(arr); i++ {
		if max < arr[i] { //找到真正的最大值
			max = arr[i]
			maxIndex = i
		}
	}
	//交换
	if maxIndex != 3 {
		arr[3], arr[maxIndex] = arr[maxIndex], arr[3]
	}

	fmt.Println("第4次 ", *arr)*/
}

func main() {
	//定义一个数组 , 从大到小
	//arr := [6]int{10, 34, 19, 100, 80, 789}

	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(900000)
	}

	//fmt.Println(arr)
	start := time.Now().Unix()
	SelectSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("选择排序耗时=%d秒", end - start)
	fmt.Println("main函数")
	//fmt.Println(arr)
}