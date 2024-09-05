package main

import (
	"fmt"
	"strings"
)

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

type ContentSegment struct {
	StartTime float32 `json:"start_time,omitempty"`
	EndTime   float32 `json:"end_time,omitempty"`
	Text      string  `json:"text,omitempty"`
}

// MatchSubtitle 给定一个音频，还有这个音频中每句话的时间段segs，保证时间段之间不会在时间轴上重叠。
// 现在给你一组同样是这个音频的时间段words，每个时间段都有一个字Text，
// 你来实现：将每个字按照StartTime升序的顺序放到 word的start_time和end_time 都在 seg的start_time和end_time之间，返回填充好Text的segs，
func MatchSubtitle(segs, words []*ContentSegment) []*ContentSegment {

	// 存储 segs 中每个时间段的索引和text列表
	segMap := make(map[int][]string)

	for _, word := range words {
		// 二分 找到 word 应该匹配到的seg 的索引
		idx := binarySearch(segs, word.StartTime)
		// 找到对应的seg, 则将word的text添加到 segMap 的text列表中
		if idx != -1 {
			segMap[idx] = append(segMap[idx], word.Text)
		}
	}

	// 遍历 segs，填充 segMap 中的text列表到 segs
	for i, seg := range segs {
		if texts, ok := segMap[i]; ok {
			seg.Text = strings.Join(texts, "")
		}
	}
	return segs
}

// 二分查找 在 segs 中找到第一个 startTime <= target <= endTime 的 seg 的索引
func binarySearch(segs []*ContentSegment, startTime float32) int {
	left, right := 0, len(segs)-1
	for left <= right {
		mid := left + (right-left)/2
		if segs[mid].StartTime <= startTime && startTime <= segs[mid].EndTime {
			return mid
		} else if startTime < segs[mid].StartTime {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	segs := []*ContentSegment{
		{StartTime: 0, EndTime: 5, Text: ""},
		{StartTime: 5, EndTime: 10, Text: ""},
		{StartTime: 10, EndTime: 15, Text: ""},
	}
	words := []*ContentSegment{
		{StartTime: 1, EndTime: 3, Text: "hello"},
		{StartTime: 6, EndTime: 8, Text: "world"},
		{StartTime: 11, EndTime: 13, Text: "golang"},
	}
	// 调用 MatchSubtitle 函数进行匹配
	result := MatchSubtitle(segs, words)

	// 输出结果
	for _, seg := range result {
		fmt.Printf("StartTime: %f, EndTime: %f, Text: %s\n", seg.StartTime, seg.EndTime, seg.Text)
	}
}
