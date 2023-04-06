package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func main() {
	type imageResp struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Data    struct {
			Size       int    `json:"size"`
			Width      int    `json:"width"`
			MaterialId int    `json:"material_id"`
			Format     string `json:"format"`
			Url        string `json:"url"`
			Signature  string `json:"signature"`
			Id         string `json:"id"`
			Height     int    `json:"height"`
		} `json:"data"`
	}

	var (
		openApiUrlPrefix = "https://ad.oceanengine.com/open_api/2/"
		uri              = "file/image/ad/"
		// 请求Header
		contentType = "multipart/form-data"
		accessToken = "e88f206ab28a97ef494b853982d81739b81a1e37"
		//XDebugMode  int = 1
		// 请求参数
		advertiserId   int64 = 1760312309087432 // 广告主ID
		uploadType           = "UPLOAD_BY_FILE" // 视频上传方式，可选值:UPLOAD_BY_FILE: 文件上传（默认值），UPLOAD_BY_URL: 网址上传
		imageSignature       = "1faaf9020e0df18fdf0429e0db211f37"
		filename             = "auto4_15151.aaaaaaa_test环境slicess__V_ZJR_ZJR_en+de_1X1_31s"                  // 素材的文件名，可自定义素材名，不传择默认取文件名，最长255个字符。UPLOAD_BY_URL必填  注：若同一素材已进行上传，重新上传不会改名。
		imageUrl             = "https://ark-oss.bettagames.com/2023-03/1faaf9020e0df18fdf0429e0db211f37.png" // 图片url地址
		//
		ttImageResp imageResp
	)
	url := fmt.Sprintf("%s%s", openApiUrlPrefix, uri)
	fileBytes, err := getFileBytes(imageUrl)
	if err != nil {
		fmt.Println("getFileBytes err", err)
		return
	}

	httpStatus, resp, err := HttpPostMultipart(url,
		map[string]string{
			"advertiser_id":   fmt.Sprintf("%d", advertiserId),
			"upload_type":     uploadType,
			"image_signature": imageSignature,
			"filename":        filename,
		},
		map[string]FileObject{
			"image_file": {
				Name:    filename,
				Content: fileBytes,
			},
		},
		map[string]string{
			"Content-Type": contentType,
			"Access-Token": accessToken,
		})
	if err != nil {
		fmt.Println("Post err", err)
		return
	}
	fmt.Println("code:", httpStatus)

	err = json.Unmarshal(resp, &ttImageResp)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	if httpStatus != http.StatusOK {
		fmt.Println("resp.StatusCode() != http.StatusOK")
		return
	}
	fmt.Println("ttImageResp: ", ttImageResp)
}

//HttpPostMultipart 通过multipart/form-data请求数据
func HttpPostMultipart(url string, formData map[string]string, fileData map[string]FileObject,
	header map[string]string) (httpStatus int, resp []byte, err error) {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for k, v := range formData {
		_ = writer.WriteField(k, v)
	}
	for k, v := range fileData {
		fileWriter, err := writer.CreateFormFile(k, v.Name)
		if err != nil {
			return 0, nil, err
		}
		_, err = io.Copy(fileWriter, bytes.NewReader(v.Content))
		if err != nil {
			return 0, nil, err
		}
	}
	err = writer.Close()
	if err != nil {
		return 0, nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return 0, nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	if len(fileData) > 0 {
		req.Header.Set("Content-Type", writer.FormDataContentType())
	}
	response, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	resp, err = ioutil.ReadAll(response.Body)
	return response.StatusCode, resp, err
}

type FileObject struct {
	Name    string
	Content []byte
}

func getFileBytes(netUrl string) ([]byte, error) {
	resp, err := resty.New().R().Get(netUrl)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
