package main

import (
	"fmt"
	"time"
)

func GetDayTimestamps(dateStr string) (int64, int64, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return 0, 0, err
	}

	startTimestamp := t.UnixNano() / 1e6
	endTimestamp := startTimestamp + 24*60*60*1000 - 1

	return startTimestamp, endTimestamp, nil
}

func main() {
	date := "2024-05-01"
	startTimestamp, endTimestamp, err := GetDayTimestamps(date)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Start Timestamp:", startTimestamp)
	fmt.Println("End Timestamp:", endTimestamp)
}
