package main

import (
	"fmt"
	"gopkg.in/resty.v1"
	"net/url"
)

func main() {
	// URL字符串
	// user
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231102%2F10007348_alias_cpsvideo_user_1698918329171.txt%3FtaskId%3DOnNLbXpKRQLcqgN&dataType=alias_user"
	// order
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231102%2F10007348_alias_cpsvideo_order_1698920405212.txt%3FtaskId%3DO4jM0nylczk118F&dataType=alias_order"
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231103%2F10007348_alias_cpsvideo_order_1698977301982.txt%3FtaskId%3DjdM86excIXp8IZ9&dataType=alias_order"
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231103%2F10007348_alias_cpsvideo_order_1698978256361.txt%3FtaskId%3DLiMSF92VqFRAkU9&dataType=alias_order"
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231106%2F10007348_alias_cpsvideo_user_1699255364687.txt%3FtaskId%3DRoxEM3IolVybj0A&dataType=alias_user"
	urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231106%2F10007348_alias_cpsvideo_user_1699265409934.txt%3FtaskId%3DC3SNHmrrnvVpJNN&dataType=alias_user"

	// 对URL进行解码
	decodedURL, err := url.QueryUnescape(urlStr)
	if err != nil {
		fmt.Println("解码失败:", err)
		return
	}

	// 打印解码后的URL
	fmt.Println("解码后的URL:", decodedURL)
	// user 解码后的URL: https://xcxoss.dzyds.com/download/channel/20231102/10007348_alias_cpsvideo_user_1698918329171.txt?taskId=OnNLbXpKRQLcqgN&dataType=alias_user
	// order 解码后的URL: https://xcxoss.dzyds.com/download/channel/20231102/10007348_alias_cpsvideo_order_1698920405212.txt?taskId=O4jM0nylczk118F&dataType=alias_order
	// order 解码后的URL: https://xcxoss.dzyds.com/download/channel/20231103/10007348_alias_cpsvideo_order_1698977301982.txt?taskId=jdM86excIXp8IZ9&dataType=alias_order
	// order 解码后的URL: https://xcxoss.dzyds.com/download/channel/20231103/10007348_alias_cpsvideo_order_1698978256361.txt?taskId=LiMSF92VqFRAkU9&dataType=alias_order
	//
	// user 解码后的URL: https://xcxoss.dzyds.com/download/channel/20231106/10007348_alias_cpsvideo_user_1699265409934.txt?taskId=C3SNHmrrnvVpJNN&dataType=alias_user
	resp, err := resty.New().SetRetryCount(3).R().
		SetHeader("Content-Type", "application/json").
		Get(decodedURL)
	if err != nil {
		fmt.Println("Get err", err)
		return
	}
	fmt.Println("resp:", resp)
}
