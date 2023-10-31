package main

import (
	"fmt"
	"time"
)

func main() {
	//loa := time.FixedZone("Asia/Shanghai", 8*60*60)
	//startDate, err := time.ParseInLocation("2006-01-02 15:04:05", "2023-10-31 01:02:36", loa)

	a := time.Duration(int64(1698684471)+7776000-time.Now().Unix()) * time.Second
	fmt.Println(a)
	if a < 0 {
		fmt.Println("小于0")
	}
}
