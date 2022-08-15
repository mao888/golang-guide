/**
    @author:胡超
    @data:2022/3/5
    @note: golang实现十大经典算法:快速排序
**/
package main

func main(){
	l := []int{0,0,0,14124,124124,1}
	quicksort(l,0,len(l)-1)
}

func partion(a []int,start int ,end int) int {
	pivot,i,j := a[start],start,end
	//将第一个元素作为pivot
	for i<j{
		for i<j&&pivot<a[j]{//从右边开始找出小于pivot的
			j--
		}
		if i<j{
			a[i] = a[j]
			i++
		}
		for i<j&&a[i]<pivot{//找出大于pivot
			i++
		}
		if i<j{
			a[j] = a[i]
			j--
		}
	}
	a[i] = pivot//这时i是pivot最终位置
	return i
}
func quicksort(a []int,start int,end int)  {
	if start<end{
		d := partion(a,start,end)
		quicksort(a,start,d-1)
		quicksort(a,d+1,end)
	}
}