package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	url := "https://ad.oceanengine.com/open_api/2/file/video/ad/"
	method := "POST"

	// 创建multipart.Writer，用于构造multipart/form-data格式的请求体
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	// 添加其他表单字段
	_ = writer.WriteField("advertiser_id", "1760312309087432")
	_ = writer.WriteField("upload_type", "UPLOAD_BY_FILE")
	_ = writer.WriteField("video_signature", "6b12a8bbbe8e69a2ef5929028b0b50c3")

	file, errFile4 := os.Open("/Users/betta/Desktop/常用/6b12a8bbbe8e69a2ef5929028b0b50c3.mp4")
	defer file.Close()

	// 创建一个multipart.Part，用于表示文件字段
	part4, errFile4 := writer.CreateFormFile("video_file", filepath.Base("/Users/betta/Desktop/常用/6b12a8bbbe8e69a2ef5929028b0b50c3.mp4"))
	// 将文件内容复制到multipart.Part中
	_, errFile4 = io.Copy(part4, file)
	if errFile4 != nil {
		fmt.Println(errFile4)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建HTTP请求
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	// 设置Content-Type为multipart/form-data
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Add("Access-Token", "e88f206ab28a97ef494b853982d81739b81a1e37")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	// 处理响应
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
