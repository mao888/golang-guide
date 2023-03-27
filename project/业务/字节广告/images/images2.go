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

	url := "https://ad.oceanengine.com/open_api/2/file/image/ad/"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("advertiser_id", "1760312309087432")
	_ = writer.WriteField("upload_type", "UPLOAD_BY_FILE")
	_ = writer.WriteField("image_signature", "1faaf9020e0df18fdf0429e0db211f37")
	file, errFile4 := os.Open("/Users/betta/Desktop/常用/1faaf9020e0df18fdf0429e0db211f37.png")
	defer file.Close()
	part4,
		errFile4 := writer.CreateFormFile("image_file", filepath.Base("/Users/betta/Desktop/常用/1faaf9020e0df18fdf0429e0db211f37.png"))
	_, errFile4 = io.Copy(part4, file)
	if errFile4 != nil {
		fmt.Println(errFile4)
		return
	}
	_ = writer.WriteField("filename", "auto4_huchao.1faaf9020e0df18fdf0429e0db211f37_test环境slicess_卡通_P_HC_HC_en_1X1_0s")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Add("Access-Token", "e88f206ab28a97ef494b853982d81739b81a1e37")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
