package sdk_release

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/mao-gutils/constants"
	gutil "github.com/mao888/mao-gutils/json"
	"github.com/mao888/mao-gutils/version"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type MProjectversions struct {
	ID          string `bson:"_id" json:"_id"`
	VersionName string `bson:"version_name" json:"version_name"`
	Status      int32  `bson:"status" json:"status"`
	Desc        string `bson:"desc" json:"desc"`
	ProjectID   int32  `bson:"project_id" json:"project_id"`
	SdkModel    string `bson:"sdk_model" json:"sdk_model"`
	Reason      string `bson:"reason" json:"reason"`
	CreateUser  string `bson:"create_user" json:"create_user"`
	SdkConfig   struct {
		Unity []struct {
			MustDown    bool   `bson:"must_down" json:"must_down"`
			Name        string `bson:"name" json:"name"`
			XmlCode     string `bson:"xml_code" json:"xml_code"`
			SdkName     string `bson:"sdk_name" json:"sdk_name"`
			SdkId       int32  `bson:"sdk_id" json:"sdk_id"`
			SdkType     string `bson:"sdk_type" json:"sdk_type"`
			VersionId   string `bson:"version_id" json:"version_id"`
			TagName     string `bson:"tag_name" json:"tag_name"`
			DownloadUrl string `bson:"download_url" json:"download_url"`
			Desc        string `bson:"desc" json:"desc"`
		} `bson:"unity" json:"unity,omitempty"`
		Ios []struct {
			MustDown    bool   `bson:"must_down" json:"must_down"`
			Name        string `bson:"name" json:"name"`
			XmlCode     string `bson:"xml_code" json:"xml_code"`
			SdkName     string `bson:"sdk_name" json:"sdk_name"`
			SdkId       int32  `bson:"sdk_id" json:"sdk_id"`
			SdkType     string `bson:"sdk_type" json:"sdk_type"`
			VersionId   string `bson:"version_id" json:"version_id"`
			TagName     string `bson:"tag_name" json:"tag_name"`
			DownloadUrl string `bson:"download_url" json:"download_url"`
			Desc        string `bson:"desc" json:"desc"`
		} `bson:"ios" json:"ios,omitempty"`
		Android []struct {
			MustDown    bool   `bson:"must_down" json:"must_down"`
			Name        string `bson:"name" json:"name"`
			XmlCode     string `bson:"xml_code" json:"xml_code"`
			SdkName     string `bson:"sdk_name" json:"sdk_name"`
			SdkId       int32  `bson:"sdk_id" json:"sdk_id"`
			SdkType     string `bson:"sdk_type" json:"sdk_type"`
			VersionId   string `bson:"version_id" json:"version_id"`
			TagName     string `bson:"tag_name" json:"tag_name"`
			DownloadUrl string `bson:"download_url" json:"download_url"`
			Desc        string `bson:"desc" json:"desc"`
		} `bson:"android" json:"android,omitempty"`
	} `bson:"sdk_config" json:"sdk_config"`
	PublishTime *time.Time `bson:"publish_time" json:"publish_time"`
	CreateTime  *time.Time `bson:"create_time" json:"create_time"`
	UpdateTime  *time.Time `bson:"update_time" json:"update_time"`
}

type SdkReleaseRecord struct {
	ID              int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SdkProjectID    int32  `gorm:"column:sdk_project_id;not null" json:"sdk_project_id"`     // SDK项目id
	VersionNumber   string `gorm:"column:version_number;not null" json:"version_number"`     // 版本号
	VersionOrdinal  string `gorm:"column:version_ordinal;not null" json:"version_ordinal"`   // 可比较的版本号
	Status          int32  `gorm:"column:status;not null" json:"status"`                     // 状态（1：Alpha版、2:Beta版、3:正式版、4:废弃）
	Model           string `gorm:"column:model;not null" json:"model"`                       // 模型（Unitypackage）
	AbandonedReason string `gorm:"column:abandoned_reason;not null" json:"abandoned_reason"` // 废弃原因
	ReleaseLog      string `gorm:"column:release_log;not null" json:"release_log"`           // 发版日志
	DocumentLink    string `gorm:"column:document_link;not null" json:"document_link"`       // 文档链接
	Android         string `gorm:"column:android;not null" json:"android"`                   // 安卓子sdk
	Ios             string `gorm:"column:ios;not null" json:"ios"`                           // ios子sdk
	Unity           string `gorm:"column:unity;not null" json:"unity"`                       // unity子sdk
	CreatorID       int32  `gorm:"column:creator_id;not null" json:"creator_id"`             // 创建人
	PublisherName   string `gorm:"column:publisher_name;not null" json:"publisher_name"`     // 发版人名称
	CreatedAt       int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`       // 发版时间
	UpdatedAt       int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`       // 更新时间
	IsDeleted       bool   `gorm:"column:is_deleted;not null" json:"is_deleted"`             // 是否删除
}

type (
	Unity struct {
		ChildSdkID    int32  `json:"child_sdk_id"`   // 子sdk项目id
		Name          string `json:"child_sdk_name"` // 子sdk名称
		VersionNumber string `json:"version_number"` // 版本号
		DownloadLink  string `json:"download_link"`  // 下载链接(手输入，则无)
	}
	Ios struct {
		ChildSdkID    int32  `json:"child_sdk_id"`   // 子sdk项目id
		Name          string `json:"child_sdk_name"` // 子sdk名称
		VersionNumber string `json:"version_number"` // 版本号
		DownloadLink  string `json:"download_link"`  // 下载链接(手输入，则无)
	}
	Android struct {
		ChildSdkID    int32  `json:"child_sdk_id"`   // 子sdk项目id
		Name          string `json:"child_sdk_name"` // 子sdk名称
		VersionNumber string `json:"version_number"` // 版本号
		DownloadLink  string `json:"download_link"`  // 下载链接(手输入，则无)
	}
)

func RunSdkReleaseRecord() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("projectversions")

	// 2、从mongo查询数据
	mProjectversions := make([]*MProjectversions, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mProjectversions)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	//fmt.Println(mProjectversions)

	// 3、将mongo数据装入切片
	sdkReleaseRecord := make([]*SdkReleaseRecord, 0)
	for i, projectversion := range mProjectversions {
		fmt.Println("sdk_release_record: ", i)
		// status
		if projectversion.Status == -1 {
			projectversion.Status = 4
		} else {
			projectversion.Status = projectversion.Status + 1
		}
		// model
		if projectversion.SdkModel == "unity" {
			projectversion.SdkModel = "Unitypackage"
		} else if projectversion.SdkModel == "android" {
			projectversion.SdkModel = "Android原生"
		} else if projectversion.SdkModel == "ios" {
			projectversion.SdkModel = "iOS原生"
		}
		// Android
		var android []Android
		for _, s := range projectversion.SdkConfig.Android {
			var an Android

			an.ChildSdkID = s.SdkId
			an.Name = s.SdkName
			an.VersionNumber = s.TagName
			an.DownloadLink = s.DownloadUrl
			// 手输
			if s.SdkId == constants.NumberZero {
				an.Name = s.Name
			}
			android = append(android, an)
		}
		androidJson, err := gutil.Object2JSONE(&android)
		if err != nil {
			return
		}
		// Ios
		var ios []Ios
		for _, s := range projectversion.SdkConfig.Ios {
			var an Ios

			an.ChildSdkID = s.SdkId
			an.Name = s.SdkName
			an.VersionNumber = s.TagName
			an.DownloadLink = s.DownloadUrl
			// 手输
			if s.SdkId == constants.NumberZero {
				an.Name = s.Name
			}
			ios = append(ios, an)
		}
		iosJson, err := gutil.Object2JSONE(&ios)
		if err != nil {
			return
		}
		// Unity
		var unity []Unity
		for _, s := range projectversion.SdkConfig.Unity {
			an := Unity{
				ChildSdkID:    s.SdkId,
				Name:          s.SdkName,
				VersionNumber: s.TagName,
				DownloadLink:  s.DownloadUrl,
			}
			unity = append(unity, an)
		}
		unityJson, err := gutil.Object2JSONE(&unity)
		if err != nil {
			return
		}

		sdkRelease := &SdkReleaseRecord{
			//ID:              0,
			SdkProjectID:    projectversion.ProjectID,
			VersionNumber:   projectversion.VersionName,
			VersionOrdinal:  version.VersionOrdinal(projectversion.VersionName),
			Status:          projectversion.Status,
			Model:           projectversion.SdkModel,
			AbandonedReason: projectversion.Reason,
			ReleaseLog:      projectversion.Desc,
			//DocumentLink:    "",
			Android:       androidJson,
			Ios:           iosJson,
			Unity:         unityJson,
			PublisherName: projectversion.CreateUser,
			CreatedAt:     projectversion.CreateTime.Unix(),
			UpdatedAt:     projectversion.UpdateTime.Unix(),
		}
		sdkReleaseRecord = append(sdkReleaseRecord, sdkRelease)
	}

	// 4、将装有mongo数据的切片入库
	err = db2.MySQLClientAdmin.Table("sdk_release_record").CreateInBatches(sdkReleaseRecord, len(sdkReleaseRecord)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
