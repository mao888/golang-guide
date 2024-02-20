package main

import (
	"fmt"
	"time"
)

func extractDates(dates []string) []string {
	if len(dates) == 0 {
		return nil
	}

	// 解析第一个日期
	firstDate, err := time.Parse("2006-01-02", dates[0])
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil
	}

	// 解析最后一个日期
	lastDate, err := time.Parse("2006-01-02", dates[len(dates)-1])
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil
	}

	// 判断是否跨月
	if firstDate.Month() == lastDate.Month() {
		// 没有跨月，返回第一天和最后一天
		return []string{lastDate.Format("2006-01-02")}
	} else {
		// 跨月了，直接返回最后一天
		return []string{firstDate.Format("2006-01-02"), lastDate.Format("2006-01-02")}
	}
}

func main() {
	dates := []string{"2023-12-30", "2023-12-31", "2024-01-01"}
	result := extractDates(dates)
	fmt.Println(result)
}
