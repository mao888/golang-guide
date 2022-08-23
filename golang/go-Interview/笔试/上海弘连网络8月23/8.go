package main

import "fmt"

func main() {
	var peoples map[string]string = make(map[string]string)
	peoples["001"] = "Jack"
	peoples["004"] = "John"
	peoples["003"] = "Georgia"
	peoples["002"] = "Lucy"
	for k, v := range peoples {
		fmt.Println(k, v)
	}
	//001 Jack
	//004 John
	//003 Georgia
	//002 Lucy
}
