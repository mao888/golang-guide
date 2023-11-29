package main

import "strconv"

func main() {
	eventParam, err := strconv.Atoi("2.1")
	if err != nil {
		panic(err)
	}
	println(eventParam)
}
