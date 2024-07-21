package main

import (
	"fmt"
	"time"
)

// GetNextFridayUTCCountdown 计算下一个 UTC 时间每周五 0 点的倒计时秒数，并返回当前是哪一周
func GetNextFridayUTCCountdown() (int64, int) {
	now := time.Now().UTC()
	_, week := now.ISOWeek()

	// 计算本周五的日期
	var nextFriday time.Time
	if now.Weekday() <= time.Friday {
		nextFriday = time.Date(now.Year(), now.Month(), now.Day()+(5-int(now.Weekday())), 0, 0, 0, 0, time.UTC)
	} else {
		// 如果今天是周五之后的某一天，计算下周五的日期
		nextFriday = time.Date(now.Year(), now.Month(), now.Day()+(7-int(now.Weekday())+5), 0, 0, 0, 0, time.UTC)
		// 因为我们计算的是下周五，所以周数需要加1
		_, week = nextFriday.ISOWeek()
	}

	// 计算倒计时秒数
	countdown := nextFriday.Sub(now).Seconds()
	return int64(countdown), week
}

func main() {
	countdown, week := GetNextFridayUTCCountdown()
	fmt.Printf("距离下一个 UTC 时间每周五 0 点的倒计时秒数: %d 秒\n", countdown)
	fmt.Printf("当前是第 %d 周\n", week)

	// 计算并显示北京时间
	now := time.Now().UTC()
	beijingTime := now.Add(8 * time.Hour)
	fmt.Printf("当前北京时间: %s\n", beijingTime.Format("2006-01-02 15:04:05"))
}
