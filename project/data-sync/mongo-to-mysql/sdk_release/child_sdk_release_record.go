package sdk_release

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// Msdkversions From Mongo
type Msdkversions struct {
	ID          string     `bson:"_id" json:"_id"`
	JenkinsName string     `bson:"jenkins_name" json:"jenkins_name"`
	TagName     string     `bson:"tag_name" json:"tag_name"`
	DownloadUrl string     `bson:"download_url" json:"download_url"`
	Desc        string     `bson:"desc" json:"desc"`
	Status      int32      `bson:"status" json:"status"`
	Unread      bool       `bson:"unread" json:"unread"`
	ProjectId   int32      `bson:"project_id" json:"project_id"`
	SdkId       int32      `bson:"sdk_id" json:"sdk_id"`
	SdkType     string     `bson:"sdk_type" json:"sdk_type"`
	PackageTime *time.Time `bson:"package_time" json:"package_time"`
	CreateTime  *time.Time `bson:"create_time" json:"create_time"`
}

// ChildSdkReleaseRecord From admin_console
type ChildSdkReleaseRecord struct {
	ID            int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ChildSdkID    int32  `gorm:"column:child_sdk_id;not null" json:"child_sdk_id"`     // 子sdk项目id
	Jenkins       string `gorm:"column:jenkins;not null" json:"jenkins"`               // Jenkins
	VersionNumber string `gorm:"column:version_number;not null" json:"version_number"` // 版本号
	DownloadLink  string `gorm:"column:download_link;not null" json:"download_link"`   // 下载链接
	PackingTime   int32  `gorm:"column:packing_time;not null" json:"packing_time"`     // 打包时间
	Remark        string `gorm:"column:remark;not null" json:"remark"`                 // 备注
	CreatedAt     int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`   // 创建时间
	UpdatedAt     int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`   // 更新时间
	IsDeleted     bool   `gorm:"column:is_deleted;not null" json:"is_deleted"`         // 是否删除
}

func RunChildSdkReleaseRecord() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("sdkversions")

	// 2、从mongo查询数据
	mSdkversions := make([]*Msdkversions, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mSdkversions)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	//fmt.Println(mSdkversions)

	// 3、将mongo数据装入切片
	childSdkReleaseRecord := make([]*ChildSdkReleaseRecord, 0)
	for i, record := range mSdkversions {
		fmt.Println("child_sdk_release_record: ", i)
		childSdk := &ChildSdkReleaseRecord{
			ChildSdkID:    record.SdkId,
			Jenkins:       record.JenkinsName,
			VersionNumber: record.TagName,
			DownloadLink:  record.DownloadUrl,
			PackingTime:   int32(record.PackageTime.Unix()),
			Remark:        record.Desc,
			CreatedAt:     record.CreateTime.Unix(),
		}
		childSdkReleaseRecord = append(childSdkReleaseRecord, childSdk)
	}

	// 4、将装有mongo数据的切片入库
	err = db2.MySQLClientAdmin.Table("child_sdk_release_record").CreateInBatches(childSdkReleaseRecord, len(childSdkReleaseRecord)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
