package main

import (
	"fmt"
	"time"
)

func main() {
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		// 兜底：用 UTC 或本地
		loc = time.Local
	}
	t := time.Unix(1774627424, 0).In(loc)
	fmt.Println(t.Format("01-02 15:04"))
}
