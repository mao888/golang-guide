package main

import "fmt"

func main() {
	c := make(chan string, 3)
	c <- "1023"
	c <- "1024"
	c <- "1025"
	close(c) // 关闭通道

	for i := 0; i < len(c); i++ {
		v, ok := <-c
		if !ok {
			fmt.Println("通道已关闭")
		} else {
			//for循环value:  1023
			//for循环value:  1024
			fmt.Println("for循环value: ", v)
		}
	}

	//for msg := range c {
	//	//range循环value:  1023
	//	//range循环value:  1024
	//	//range循环value:  1025
	//	fmt.Println("range循环value: ", msg)
	//}

	//for {
	//	v, ok := <-c
	//	if !ok {
	//		fmt.Println("通道已关闭")
	//		break
	//	}
	//	//for-select循环value:  1023
	//	//for-select循环value:  1024
	//	//for-select循环value:  1025
	//	fmt.Println("for-select循环value: ", v)
	//}
}
