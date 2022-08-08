/**
    @author:Hasee
    @data:2022/3/14
    @note:
**/
package main

func main()  {
	add([]int{1,2}...)
}

func add(args ...int) int {
	sum := 0
	for _,arg := range args{
		sum += arg
	}
	return sum
}
