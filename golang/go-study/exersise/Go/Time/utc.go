package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UTC().Format("2006-01-02 15:04:05"))
	sub := time.Now().UTC().Sub(time.Now())

	fmt.Println(sub)
}
