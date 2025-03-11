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
	// 示例时间戳（2023-10-05 12:00:00 UTC）
	timestamp := int64(1733906270)

	// 获取 GMT+8 的时区
	location := GetGMTLocation(5.5)

	// 获取周
	week := GetWeek(timestamp, location)
	fmt.Printf("周：%s\n", week)

	// 获取月
	month := GetMonth(timestamp, location)
	fmt.Printf("月：%s\n", month)
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
