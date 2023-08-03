package main

import (
	"fmt"
	"time"
)

func main() {
	chuo := int64(1688589582)
	//bendi := time.Unix(int64(chuo), 0).Format("2006-01-02 15:04:05")
	//lastLoginAt, err := time.Parse("2006-01-02 15:04:05", bendi)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(lastLoginAt.String())

	// 将时间戳转换为北京时间
	beijng := time.Unix(chuo, 0).In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04:05")

	// 将时间戳转换为UTC时间
	utc := time.Unix(chuo, 0).UTC().Format("2006-01-02 15:04:05")

	fmt.Println("beijng", beijng)
	fmt.Println("utc", utc)
}
