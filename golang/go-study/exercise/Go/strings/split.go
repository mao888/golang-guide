package main

import (
	"fmt"
	"strings"
)

// 模拟 req.UrlGroup
type Request struct {
	UrlGroup []struct {
		FileName string
	}
}

func main() {
	// 模拟请求
	req := Request{
		UrlGroup: []struct {
			FileName string
		}{
			{FileName: "auto4_8100461.222222222_test环境slicess_非原始_V_HC+TZH_HC_en_4X5_36s.mp4"},
			{FileName: "auto4_8100461.222222222_test环境slicess_非原始_V_HC+TZH_HC_en_4X5_36s.mp4"},
			{FileName: "auto4_8100461.222222222_test环境slicess_非原始_V_HC+TZH_HC_en_4X5_37s.mp4"},
			{FileName: "auto4_8100461.222222222_test环境slicess_非原始_V_HC+TZH_HC_en_4X5_38s.mp4"},
			{FileName: "auto4_8100461.222222222_test环境slicess_非原始_V_HC+TZH_HC_en_4X5_38s.mp4"},
			{FileName: "auto4_8100461.222222222_test环境slicess_非原始_V_HC+TZH_HC_en_4X5_37s.mp4"},
			{FileName: "auto4_8100461.222222222_test环境slicess_非原始_V_HC+TZH_HC_en_4X5_37s.mp4"},
			{FileName: "auto4_8100461.222222222_test环境slicess_非原始_V_HC+TZH_HC_en_4X5_36s.mp4"},
		},
	}

	// 遍历 req.UrlGroup，如果 FileName 重复，则名字添加递增数字
	counts := make(map[string]int)

	for i, u := range req.UrlGroup {
		if count, ok := counts[u.FileName]; ok {
			req.UrlGroup[i].FileName = addSuffix(u.FileName, count)
			counts[u.FileName]++
		} else {
			counts[u.FileName] = 1
		}
	}

	// 打印修改后的文件名
	for _, u := range req.UrlGroup {
		fmt.Println(u.FileName)
	}
}

func addSuffix(fileName string, count int) string {
	parts := strings.Split(fileName, ".")
	baseName := strings.Join(parts[:len(parts)-1], ".")
	extension := parts[len(parts)-1]
	return fmt.Sprintf("%s_%d.%s", baseName, count+1, extension)
}
