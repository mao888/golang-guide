package main

import (
	"fmt"
	"time"
)

//func GetGMTLocation(h int) *time.Location {
//	if h > 0 {
//		return time.FixedZone(fmt.Sprintf("GMT+%d", h), h*60*60)
//	}
//	return time.FixedZone(fmt.Sprintf("GMT%d", h), h*60*60)
//}

// GetGMTLocation 根据与 GMT 的时差（小时，支持浮点数）返回对应的时区
func GetGMTLocation(h float64) *time.Location {
	// 计算时区偏移量（秒）
	offset := int(h * 60 * 60)
	if h > 0 {
		return time.FixedZone(fmt.Sprintf("GMT+%.1f", h), offset)
	}
	return time.FixedZone(fmt.Sprintf("GMT%.1f", h), offset)
}

func main() {
	// 示例时间戳
	timestamp := int64(1741586735)

	// 获取 GMT+8 的时区
	location := GetGMTLocation(8.0)

	// 获取周
	week := GetWeek(timestamp, location)
	fmt.Printf("周：%s\n", week)

	// 获取月
	month := GetMonth(timestamp, location)
	fmt.Printf("月：%s\n", month)

	// 获取下个月 0 点的时间戳
	nextMonthTimestamp := GetNextMonthStartTimestamp(timestamp, location)
	fmt.Printf("下个月 0 点的时间戳：%d\n", nextMonthTimestamp)

	// 获取下周 0 点的时间戳
	nextWeekTimestamp := GetNextWeekStartTimestamp(timestamp, location)
	fmt.Printf("下周 0 点的时间戳：%d\n", nextWeekTimestamp)

	// 获取当前时间戳
	currentTimestamp := time.Now().Unix()

	// 月倒计时
	countdown := nextMonthTimestamp - currentTimestamp
	fmt.Printf("月倒计时：%d\n", countdown)

	// 周倒计时
	countdownWeek := nextWeekTimestamp - currentTimestamp
	fmt.Printf("周倒计时：%d\n", countdownWeek)
}

// GetWeek 根据时间戳和时区返回周（格式 202503）
func GetWeek(timestamp int64, location *time.Location) string {
	// 将时间戳转换为时间对象
	t := time.Unix(timestamp, 0).In(location)

	// 获取年份和周数
	year, week := t.ISOWeek()

	// 格式化周（例如 202503）
	return fmt.Sprintf("%04d%02d", year, week)
}

// GetMonth 根据时间戳和时区返回月（格式 202503）
func GetMonth(timestamp int64, location *time.Location) string {
	// 将时间戳转换为时间对象
	t := time.Unix(timestamp, 0).In(location)

	// 获取年份和月份
	year, month := t.Year(), int(t.Month())

	// 格式化月（例如 202503）
	return fmt.Sprintf("%04d%02d", year, month)
}

// GetNextMonthStartTimestamp 获取下个月 0 点的时间戳
func GetNextMonthStartTimestamp(timestamp int64, location *time.Location) int64 {
	// 将时间戳转换为时间对象
	t := time.Unix(timestamp, 0).In(location)

	// 计算下个月的 0 点时间
	nextMonth := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, location)

	// 返回时间戳
	return nextMonth.Unix()
}

// GetNextWeekStartTimestamp 获取下周 0 点的时间戳
func GetNextWeekStartTimestamp(timestamp int64, location *time.Location) int64 {
	// 将时间戳转换为时间对象
	t := time.Unix(timestamp, 0).In(location)
	fmt.Printf("t: %v\n", t)

	// 计算下周的 0 点时间
	daysUntilNextWeek := (7 - int(t.Weekday())) % 7 // 距离下周的天数
	nextWeek := time.Date(t.Year(), t.Month(), t.Day()+daysUntilNextWeek+1, 0, 0, 0, 0, location)

	fmt.Printf("nextWeek: %v\n", nextWeek)
	// 返回时间戳
	return nextWeek.Unix()
}
