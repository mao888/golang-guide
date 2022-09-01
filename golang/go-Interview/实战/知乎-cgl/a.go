package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s) //	[0 0 0 0 0 1 2 3]

	s2 := make([]int, 0)
	s2 = append(s2, 1, 2, 3, 4)
	fmt.Println(s2) //	[1 2 3 4]

	a := [3]int{7, 8, 9}
	//a := make([]int, 1)
	//fmt.Printf("%+v\n", a) //	[7 8 9]
	//b := ap(a)
	//ap(a)
	//fmt.Printf("%+v\n", a) //	[7 8 9]
	app(a)
	fmt.Printf("%+v\n", a) //	[1 8 9]
}

func ap(a []int) {
	a = append(a, 10)
}

func ap1(a []int) []int {
	a = append(a, 10)
	return a
}

func app(a [3]int) {
	a[0] = 1
}
