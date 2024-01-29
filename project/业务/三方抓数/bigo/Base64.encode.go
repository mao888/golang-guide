package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	clientID := "2401244203476908609792"
	secret := "chv65g99dlw9"

	// 拼接client_id和secret，并转换为字节数组
	data := []byte(clientID + ":" + secret)

	// 使用base64进行编码
	encoded := base64.StdEncoding.EncodeToString(data)
	authorization := "Basic " + encoded
	fmt.Println("Base64 Encoded:", authorization) // Basic MjQwMTI0NDIwMzQ3NjkwODYwOTc5MjpjaHY2NWc5OWRsdzk=
}
