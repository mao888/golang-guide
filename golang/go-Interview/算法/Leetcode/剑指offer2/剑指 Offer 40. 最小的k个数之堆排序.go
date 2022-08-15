/**
    @author:胡超
    @data:2022/3/8
    @note:剑指 Offer 40. 最小的k个数 之 堆排序
**/
package main

import "fmt"

// 堆排序
func main() {
	arr := []int{0,1,2,1}
	fmt.Println(getLeastNumbers(arr,1))
}

func getLeastNumbers(arr []int, k int) []int {
	HeapSort(arr)
	a := make([]int,0,k)
	for i :=0 ; i<k ; i++ {
		a = append(a, arr[i])
	}
	return a
}

// 编写一个堆排序的方法
func HeapSort(arr []int) []int {
	/**
	 * 1.无序序列构建成一个堆，根据升序降序需求选择大顶堆或小顶堆
	 * */
	// 构造一个大顶堆
	for i := len(arr)/2-1 ; i>=0;i--{
		adjustHeap(arr,i,len(arr))
	}
	/**
	 *2.将堆顶元素与末尾元素交换，将最大元素“沉”到数组末端
	 *3.重新调整结构，使其满足堆定义，然后继续交换堆顶元素与当前末尾元素，反复执行调整+交换步骤，直到这个序列有序
	 * */
	for j := len(arr)-1;j>0;j-- {
		arr[0],arr[j] = arr[j],arr[0]
		adjustHeap(arr,0,j)
	}
	return arr
}

// 将一个数组(二叉树)，调整成一个大顶堆
func adjustHeap(arr []int,i int, length int)  {
	temp := arr[i];		// // 先取出当前元素的值，保存在临时变量
	// 开始调整
	// 说明:k = 2*i+1 是 i 的左子结点
	for k := 2*i+1;k < length; k = 2*k+1 {
		if k+1< length && arr[k] < arr[k+1] {
			k++;
		}
		if arr[k] > temp {
			arr[i] = arr[k];
			i = k;
		}else {
			break
		}
	}
	// 当for循环结束后，我们已经将以i为父结点的树的最大值，放在了最顶(以i为父结点的局部二叉树)
	arr[i] = temp
}