/**
    @author:Hasee
    @data:2022/3/19
    @note:
**/
package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student {
	s := new(Student) //局部变量s逃逸到堆

	s.Name = name
	s.Age = age

	return s
}

func main() {
	fmt.Println(StudentRegister("Jim", 18))
}
