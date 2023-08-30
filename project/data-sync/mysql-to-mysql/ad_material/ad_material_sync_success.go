package ad_material

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// AdMaterialSyncSuccess mapped from table <ad_material_sync_success>
type AdMaterialSyncSuccess struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MaterilaType int32  `gorm:"column:materila_type;not null;default:1" json:"materila_type"` // '附件文件类型， 1: file,  2: image,3: video',
	MaterialID   int32  `gorm:"column:material_id;not null" json:"material_id"`               // 素材id
	Name         string `gorm:"column:name;not null" json:"name"`                             // 素材名称 拼接而成
	URL          string `gorm:"column:url;not null" json:"url"`                               // 素材源地址
	MaterialMd5  string `gorm:"column:material_md5;not null" json:"material_md5"`             // 素材md5
	AccountID    string `gorm:"column:account_id;not null" json:"account_id"`                 // 所属账户
	Creator      int32  `gorm:"column:creator;not null" json:"creator"`                       // 创建者
	Type         int32  `gorm:"column:type;not null;default:1" json:"type"`                   // 上传日志类型 1：Facebook 2：YouTube 3：优量汇 4：今日头条
	SuccessID    string `gorm:"column:success_id;not null" json:"success_id"`                 // fb 返回 结果id
	BatchID      string `gorm:"column:batch_id;not null" json:"batch_id"`                     // 批处理id
}

// GetVideoMaterialResp 获取视频素材应答字段
type GetVideoMaterialResp struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
	Data      struct {
		List []struct {
			BitRate          int           `json:"bit_rate"`
			CreateTime       string        `json:"create_time"`
			Duration         float64       `json:"duration"`
			Filename         string        `json:"filename"`
			Format           string        `json:"format"`
			Height           int           `json:"height"`
			Id               string        `json:"id"`
			Labels           []string      `json:"labels"`
			MaterialId       int64         `json:"material_id"`
			OrganizationTags []interface{} `json:"organization_tags"`
			PosterUrl        string        `json:"poster_url"`
			Signature        string        `json:"signature"`
			Size             int           `json:"size"`
			Source           string        `json:"source"`
			Url              string        `json:"url"`
			Width            int           `json:"width"`
		} `json:"list"`
		PageInfo struct {
			Page        int `json:"page"`
			PageSize    int `json:"page_size"`
			TotalNumber int `json:"total_number"`
			TotalPage   int `json:"total_page"`
		} `json:"page_info"`
	} `json:"data"`
}

// GetImageMaterialResp 获取图片素材应答字段
type GetImageMaterialResp struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
	Data      struct {
		List []struct {
			Aigc       bool   `json:"aigc"`
			CreateTime string `json:"create_time"`
			Filename   string `json:"filename"`
			Format     string `json:"format"`
			Height     int    `json:"height"`
			Id         string `json:"id"`
			MaterialId int64  `json:"material_id"`
			Signature  string `json:"signature"`
			Size       int    `json:"size"`
			Url        string `json:"url"`
			Width      int    `json:"width"`
		} `json:"list"`
		PageInfo struct {
			Page        int `json:"page"`
			PageSize    int `json:"page_size"`
			TotalNumber int `json:"total_number"`
			TotalPage   int `json:"total_page"`
		} `json:"page_info"`
	} `json:"data"`
}

var videoImageIdMaterialIdMap = map[string]int64{}

func RunAdMaterialSyncSuccess() {

	// 1、获取 accountID
	var accountID []string
	err := db2.MySQLClientCruiser.Table("ad_material_sync_success").Distinct("account_id").
		Where("type = ?", 4).Where("id <= ?", 729610).Group("account_id").Find(&accountID).Error
	if err != nil {
		glog.Errorf("查询错误：%s", err)
		return
	}
	// 2、获取 每个accountID下的 视频素材,图片素材
	//	并保存视频id、图片id与素材id的映射关系
	for _, v := range accountID {
		glog.Infof("accountID: %s", v)
		// 获取视频素材 保存视频id 与 素材id 映射关系
		err := GetVideoMaterial(v)
		if err != nil {
			glog.Errorf("GetVideoMaterial 获取视频素材错误：%s", err)
			return
		}
		fmt.Println("videoImageIdMaterialIdMap1: ", videoImageIdMaterialIdMap)
		// 获取图片素材 保存图片id 与 素材id 映射关系
		err = GetImageMaterial(v)
		if err != nil {
			glog.Errorf("GetImageMaterial 获取图片素材错误：%s", err)
			return
		}
		fmt.Println("videoImageIdMaterialIdMap2: ", videoImageIdMaterialIdMap)
	}
	fmt.Println("videoImageIdMaterialIdMap总: ", videoImageIdMaterialIdMap)
	// 3、根据 accountID 获取 ad_material_sync_success
	var adMaterialSyncSuccess = make([]*AdMaterialSyncSuccess, 0)
	err = db2.MySQLClientCruiser.Table("ad_material_sync_success").
		Where("account_id IN ?", accountID).Where("id <= ?", 729610).
		Find(&adMaterialSyncSuccess).Error
	if err != nil {
		glog.Errorf("查询错误：%s", err)
		return
	}

	// 3、更新 ad_material_sync_success
	var sqlListUpdate []string
	var sqlListUpdateRecover []string
	for _, adMaterialSyncSuccess := range adMaterialSyncSuccess {
		// 写出所有条目的更新sql语句, 并导出到本地文件中
		glog.Infof("success_id: %s", adMaterialSyncSuccess.SuccessID)
		sqlUpdate := fmt.Sprintf("UPDATE ad_material_sync_success SET success_id = %d WHERE id = %d;",
			videoImageIdMaterialIdMap[adMaterialSyncSuccess.SuccessID], adMaterialSyncSuccess.ID)
		sqlListUpdate = append(sqlListUpdate, sqlUpdate)
		glog.Infof("update_sql: %s", sqlUpdate)

		sqlUpdateRecover := fmt.Sprintf("UPDATE ad_material_sync_success SET success_id = %s WHERE id = %d;",
			adMaterialSyncSuccess.SuccessID, adMaterialSyncSuccess.ID)
		sqlListUpdateRecover = append(sqlListUpdateRecover, sqlUpdateRecover)
		glog.Infof("update_sql_recover: %s", sqlUpdateRecover)
	}

	// 更新sql 写出到本地文件
	err = WriteToFile(sqlListUpdate, "update_success_id.sql")
	if err != nil {
		glog.Errorf("更新sql 写出到本地文件错误：%s", err)
		return
	}
	// 恢复sql 写出到本地文件
	err = WriteToFile(sqlListUpdateRecover, "update_success_id_recover.sql")
	if err != nil {
		glog.Errorf("恢复sql 写出到本地文件错误：%s", err)
		return
	}
}

// GetVideoMaterial 获取视频素材
func GetVideoMaterial(advertiserId string) error {
	// 当total_page > 1 时，需要循环请求
	totalPage, err := getVideo(advertiserId, 1)
	if err != nil {
		glog.Errorf("getVideo 请求错误：%s", err)
		return err
	}
	if totalPage > 1 {
		// 从第二页开始请求
		for i := 2; i <= totalPage; i++ {
			_, err := getVideo(advertiserId, i)
			if err != nil {
				glog.Errorf("getVideo 请求错误：%s", err)
				return err
			}
		}
	}
	return nil
}

// getVideo 请求视频素材
func getVideo(advertiserId string, page int) (int, error) {
	url := "https://ad.oceanengine.com/open_api/2/file/video/get/"
	method := "GET"
	payload := strings.NewReader(fmt.Sprintf(`{
    			"advertiser_id": %s,
   				 "page":%d,
   				 "page_size":100}`, advertiserId, page))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		glog.Errorf("请求错误：%s", err)
		return 0, err
	}
	req.Header.Add("Access-Token", "c0bfed085c0dbc19bb2e41920ecd6d2a6d398b32")
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		glog.Errorf("请求错误 视频素材：%s", err)
		return 0, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		glog.Errorf("请求错误 视频素材：%s", err)
		return 0, err
	}
	glog.Infof("resp: %s", string(body))
	// 解析返回值
	var getVideoMaterialResp GetVideoMaterialResp
	err = json.Unmarshal(body, &getVideoMaterialResp)
	if err != nil {
		glog.Errorf("解析错误：%s", err)
		return 0, err
	}
	if getVideoMaterialResp.Code != 0 {
		glog.Errorf("请求错误：%s", getVideoMaterialResp.Message)
		return 0, err
	}
	// 保存 视频id 与 素材id 映射关系
	for _, s := range getVideoMaterialResp.Data.List {
		videoImageIdMaterialIdMap[s.Id] = s.MaterialId
	}
	return getVideoMaterialResp.Data.PageInfo.TotalPage, nil
}

// GetImageMaterial 获取图片素材
func GetImageMaterial(advertiserId string) error {
	totalPage, err := getImage(advertiserId, 1)
	if err != nil {
		glog.Errorf("getImage 请求错误：%s", err)
		return err
	}
	// 当total_page > 1 时，需要循环请求
	if totalPage > 1 {
		// 从第二页开始请求
		for i := 2; i <= totalPage; i++ {
			_, err := getImage(advertiserId, i)
			if err != nil {
				glog.Errorf("getImage 请求错误：%s", err)
				return err
			}
		}
	}
	return nil
}

// getImage 请求图片素材
func getImage(advertiserId string, page int) (int, error) {
	url := "https://api.oceanengine.com/open_api/2/file/image/get/"
	method := "GET"
	payload := strings.NewReader(fmt.Sprintf(`{
    			"advertiser_id": %s,
   				 "page":%d,
   				 "page_size":100}`, advertiserId, page))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		glog.Errorf("请求错误：%s", err)
		return 0, err
	}
	req.Header.Add("Access-Token", "c0bfed085c0dbc19bb2e41920ecd6d2a6d398b32")
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		glog.Errorf("请求错误 图片素材：%s", err)
		return 0, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		glog.Errorf("请求错误 图片素材：%s", err)
		return 0, err
	}
	fmt.Println(string(body))
	// 解析返回值
	var getImageMaterialResp GetImageMaterialResp
	err = json.Unmarshal(body, &getImageMaterialResp)
	if err != nil {
		glog.Errorf("解析错误：%s", err)
		return 0, err
	}
	if getImageMaterialResp.Code != 0 {
		glog.Errorf("请求错误：%s", getImageMaterialResp.Message)
		return 0, err
	}
	// 保存 视频id 与 素材id 映射关系
	for _, s := range getImageMaterialResp.Data.List {
		videoImageIdMaterialIdMap[s.Id] = s.MaterialId
	}
	return getImageMaterialResp.Data.PageInfo.TotalPage, nil
}

// WriteToFile 写出到本地文件 封装一个函数
func WriteToFile(sqlList []string, fileName string) error {
	// 创建目录，如果不存在
	directory := "/Users/betta/GolandProjects/my-project/golang-guide/project/data-sync/mysql-to-mysql/ad_material"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.MkdirAll(directory, 0755)
		if err != nil {
			glog.Errorf("创建目录错误：%s", err)
			return err
		}
	}

	// 将SQL语句连接起来，每条SQL占一行
	content := ""
	for _, sql := range sqlList {
		content += sql + "\n"
	}

	// 写入文件
	fullPath := filepath.Join(directory, fileName)
	err := ioutil.WriteFile(fullPath, []byte(content), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("SQL statements written to: %s\n", fullPath)
	return nil
}
