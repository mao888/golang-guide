package main

import (
	"fmt"
	"net/url"
)

func main() {
	// URL字符串
	urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231102%2F10007348_alias_cpsvideo_user_1698918329171.txt%3FtaskId%3DOnNLbXpKRQLcqgN&dataType=alias_user"

	// 对URL进行解码
	decodedURL, err := url.QueryUnescape(urlStr)
	if err != nil {
		fmt.Println("解码失败:", err)
		return
	}

	// 打印解码后的URL
	fmt.Println("解码后的URL:", decodedURL)
	// 用户数据查询 解码后的URL: https://xcxoss.dzyds.com/download/channel/20231102/10007348_alias_cpsvideo_user_1698918329171.txt?taskId=OnNLbXpKRQLcqgN&dataType=alias_user
}
