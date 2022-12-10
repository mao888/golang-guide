package main

import "fmt"

func main() {
	// 假设序列中的每一个元素表示[start, end)的时间段
	timeSeq := [][]int{{1, 9}, {2, 4}, {5, 6}, {3, 7}, {1, 8}}

	// 根据时间段中的起始时间和结束时间，创建一个新的时间序列，表示每个时间点上连接和断开的设备数量
	timePoints := make(map[int]int)
	for _, t := range timeSeq {
		timePoints[t[0]]++
		timePoints[t[1]]--
	}

	// 记录同一时间最多连接的设备数量
	maxConnectedDevices := 0
	// 记录当前时间连接的设备数量
	currentConnectedDevices := 0
	// 遍历每个时间点，计算同一时间最多连接的设备数量
	for _, v := range timePoints {
		currentConnectedDevices += v
		if currentConnectedDevices > maxConnectedDevices {
			maxConnectedDevices = currentConnectedDevices
		}
	}

	// 输出同一时间最多连接的设备数量
	fmt.Println("最多同时连接的设备数量：", maxConnectedDevices)
}
