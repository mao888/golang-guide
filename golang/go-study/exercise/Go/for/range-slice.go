package main

import "fmt"

func main() {
	words := []string{"Go", "语言", "高性能", "编程"}

	for i, s := range words {
		words = append(words, "test")
		fmt.Println(i, s)
	}
	//0 Go
	//1 语言
	//2 高性能
	//3 编程

	//for i := 0; i < len(words); i++ {
	//	words = append(words, "test")
	//	fmt.Println(i, words[i])
	//	if len(words) >= 10 {
	//		break
	//	}
	//}
	//0 Go
	//1 语言
	//2 高性能
	//3 编程
	//4 test
	//5 test

}
