package main

import (
	"fmt"
	"time"
)

func getMonthRange(dateStr string) (string, string) {
	// 解析日期字符串
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return "", ""
	}

	// 获取当月第一天
	firstDay := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())

	// 获取当月最后一天
	lastDay := firstDay.AddDate(0, 1, -1)

	return firstDay.Format("2006-01-02"), lastDay.Format("2006-01-02")
}

func main() {
	dateStr := "2023-02-11"
	firstDay, lastDay := getMonthRange(dateStr)
	fmt.Println("当月第一天:", firstDay)
	fmt.Println("当月最后一天:", lastDay)
}
