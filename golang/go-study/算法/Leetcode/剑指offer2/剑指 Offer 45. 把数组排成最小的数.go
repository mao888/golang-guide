/**
    @author:Huchao
    @data:2022/3/6
    @note: 剑指 Offer 45. 把数组排成最小的数
**/
package main

import (
	"fmt"
	"strings"
)

func main()  {
	//[3,30,34,5,9]
	arr := []int{3,30,34,5,9}
	//minNumber(arr)
	fmt.Println(minNumber(arr))
}

func minNumber(nums []int) string {
	BubbleSort(nums)
	temp := make([]string,len(nums))
	for i,num := range nums{
		temp[i] = fmt.Sprintf("%d",num)
	}
	result := strings.Join(temp,"")
	return strings.ReplaceAll(result,"\\[|]|,|\\s", "")
}

//冒泡排序
func BubbleSort(arr []int) []int {
	for i :=0;i<len(arr);i++ {
		for j:=i+1;j<len(arr);j++ {
			if strings.Compare(string(arr[i]+arr[j]),string(arr[j]+arr[i]))>0{
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}
	return arr
}