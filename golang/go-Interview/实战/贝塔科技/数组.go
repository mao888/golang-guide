package main

import "fmt"

func main() {
	args := [5]int{0, 1, 2, 3, 4}
	slice1 := args[0:2]
	fmt.Println("len(slice1):", len(slice1)) //	2
	fmt.Println("cap(slice1):", cap(slice1)) //	5
	fmt.Printf("%p\n", slice1)               // 0xc0000aa060
	slice1 = append(slice1, 5)
	fmt.Printf("%p\n", slice1)               // 0xc0000aa060
	fmt.Println(args)                        //	[0 1 5 3 4]
	fmt.Println("len(slice1):", len(slice1)) // 3
	fmt.Println("cap(slice1):", cap(slice1)) // 5
	fmt.Println(slice1)                      //	[0 1 5]
	slice1 = append(slice1, 10, 11, 12)
	fmt.Printf("%p\n", slice1)               // 0xc0000b6000
	fmt.Println(args)                        //	[0 1 5 3 4]
	fmt.Println(slice1)                      //	[0 1 5 10 11 12]
	fmt.Println("len(slice1):", len(slice1)) // 6
	fmt.Println("cap(slice1):", cap(slice1)) // 10
}
