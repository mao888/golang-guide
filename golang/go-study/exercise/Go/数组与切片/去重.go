package main

import "fmt"

func main() {
	fmt.Printf("去重之后: %v", removeRepByMap([]string{"a", "b", "c", "d", "a", "b", "c", "d"}))
}

// 切片去重
func removeRepByMap(slc []string) []string {
	result := make([]string, 0, len(slc))
	tempMap := make(map[string]struct{}, len(slc))
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = struct{}{}
		if len(tempMap) != l { // 不存在
			result = append(result, e)
		}
	}
	return result
}
