package main
import (
	"fmt"
	"math/rand"
	"time"
)

//快速排序
//说明
//1. left 表示 数组左边的下标
//2. right 表示数组右边的下标
//3  array 表示要排序的数组
func QuickSort(left int, right int, array *[8000000]int) {
	l := left
	r := right
	// pivot 是中轴， 支点
	pivot := array[(left + right) / 2]
	temp := 0

	//for 循环的目标是将比 pivot 小的数放到 左边
	//  比 pivot 大的数放到 右边
	for ; l < r; {
		//从  pivot 的左边找到大于等于pivot的值
		for ; array[l] < pivot; {
			l++
		}
		//从  pivot 的右边边找到小于等于pivot的值
		for ; array[r] > pivot; {
			r--
		}
		// 1 >= r 表明本次分解任务完成, break
		if l >= r { 
			break
		}
		//交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		//优化
		if array[l]== pivot  {
			r--
		}
		if array[r]== pivot {
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
}


func main() {

	// arr := [9]int {-9,78,0,23,-567,70, 123, 90, -23}
	// fmt.Println("初始", arr)

	var arr [8000000]int
	for i := 0; i < 8000000; i++ {
		arr[i] = rand.Intn(900000)
	}

	//fmt.Println(arr)
	start := time.Now().Unix()
	//调用快速排序
	QuickSort(0, len(arr) - 1, &arr)
	end := time.Now().Unix()
	fmt.Println("main..")
	fmt.Printf("快速排序法耗时%d秒", end - start)
	//fmt.Println(arr)

}