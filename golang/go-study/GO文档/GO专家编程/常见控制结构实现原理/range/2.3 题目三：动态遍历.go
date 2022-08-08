/**
    @author:Hasee
    @data:2022/3/18
    @note:
**/
package main

func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
}