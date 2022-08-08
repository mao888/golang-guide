package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time
	now = time.Now()
	fmt.Println(now)

	// var secs time.Time
	secs := now.Unix()
	fmt.Println(secs)
}
