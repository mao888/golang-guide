package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	// 假设 resp 是包含了你提供的文本数据的字符串
	resp := `{"xuniPay":"1","ver":"270","os":"0","openId":"o6yWE5AQwkRJJjIGAzVmkTInrBes","referralId":"13275156","discount":39.9,"referralName":"贝塔+ZTF+名门-低","moneyBenefit":"31.92","type":1,"userId":1041609501,"promotionId":"7297088096628473892","scene":"1069","sourceInfo":"41000100737","sourceDesc":"我的名门女总裁","chapterId":"10","appId":"wx0c2a9d4d32b5ceed","domain":16,"ctime":1699065007,"id":237264531,"statusNotify":0,"projectId":"7297087194023706636","dyeTime":"1699064124","channelId":"29578","registerDate":"1699064124"}
{"xuniPay":"1","openId":"o6yWE5AnQCnUn5_A9wLYPaNFkPwg","discount":9.9,"referralName":"贝塔+ZTF+名门-低","moneyBenefit":"7.92","type":1,"promotionId":"7297088096628473892","scene":"1069","sourceInfo":"41000100737","chapterId":"10","appId":"wx0c2a9d4d32b5ceed","outTradeNo":"4200001988202311048056119460","ctime":1699055823,"id":237082861,"statusNotify":1,"channelId":"29578","registerDate":"1699055123","finishTime":1699055835,"ver":"270","os":"0","referralId":"13275156","userId":1041132241,"sourceDesc":"我的名门女总裁","domain":16,"projectId":"7297087194023706636","dyeTime":"1699055123"}`

	// 分割 resp 字符串为每个 JSON 对象
	jsonObjects := strings.Split(resp, "\n")

	// 创建一个存储 JSON 对象的切片
	var jsonArray []map[string]interface{}

	// 遍历每个 JSON 对象，解析为 map 并添加到切片中
	for _, jsonObj := range jsonObjects {
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(jsonObj), &data); err == nil {
			jsonArray = append(jsonArray, data)
		}
	}

	// 打印 JSON 数组
	jsonArrayStr, _ := json.Marshal(jsonArray)
	fmt.Println(string(jsonArrayStr))

}
