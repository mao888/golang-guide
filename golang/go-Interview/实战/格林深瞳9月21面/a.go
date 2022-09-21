package main

import "fmt"

// 一个西瓜7快，苹果3块，荔枝1块3个，50块钱 买 50个水果 多少中买法
// 一个西瓜7快，苹果3块，荔枝1块3个，50块钱  多少中买法

//7x + 3y + z/3 = 50
//x + y + z = 50

func main() {
	//int x=0, y = 0, z = 0;
	var count = 0

	for x := 1; x <= 50/7; x++ {
		for y := 1; y <= 50/3; y++ {
			for z := 1; z <= 150; z++ {
				if (7*x+3*y+1/3*z == 50) && z%3 == 0 {
					count++
					fmt.Printf("%d,%d,%d", x, y, z)
					fmt.Println()
				}
			}
		}
	}
	fmt.Println(count)

}
