package main

import (
	"fmt"
	"github.com/mao888/mao-gutils/constants"
	"time"
)

func ExtractDates(dates []string) ([]string, error) {
	if len(dates) == 0 {
		return []string{}, nil
	}

	// 解析第一个日期
	firstDate, err := time.Parse(constants.TimeYMD, dates[0])
	if err != nil {
		return nil, err
	}

	// 解析最后一个日期
	lastDate, err := time.Parse(constants.TimeYMD, dates[len(dates)-1])
	if err != nil {
		return nil, err
	}

	// 判断是否跨月
	if firstDate.Month() == lastDate.Month() {
		// 没有跨月，直接返回最后一天
		return []string{dates[len(dates)-1]}, nil
	} else {
		// 跨月了，返回第一天和最后一天
		return []string{dates[0], dates[len(dates)-1]}, nil
	}
}
func main() {
	dates := []string{"2023-12-30", "2023-12-31", "2024-01-01"}
	result, err := ExtractDates(dates)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
