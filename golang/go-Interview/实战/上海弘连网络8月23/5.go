package main

import "fmt"

func main() {
	Defer("John")
}

func Defer(name string) {
	defer func(param string) {
		fmt.Printf("%s", param)
	}(name)

	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("%s", err)
		}
	}()
	name = "Lee"
	panic("error")
	defer func() {
		fmt.Printf("end") //	errorJohn
	}()
}
