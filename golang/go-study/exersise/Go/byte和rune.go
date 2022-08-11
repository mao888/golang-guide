/**
    @author: edy
    @since: 2022/8/11
    @desc: //TODO
**/
package main

import "fmt"

//https://blog.csdn.net/qq_42410605/article/details/114818366

func main() {
	s := "asc"
	fmt.Printf("%T  ", s[0])
	fmt.Println()
	for i := range s {
		fmt.Printf("%T ", s[i])
	}
	fmt.Println()
	for _, i2 := range s {
		fmt.Printf("%T ", i2)
	}
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%T ", s[i])
	}
}
