package sdk_release

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// MSdks From Mongo
type MSdks struct {
	ID          int32      `bson:"_id" json:"_id"`
	JenkinsName string     `bson:"jenkins_name" json:"jenkins_name"`
	JenkinsUrl  string     `bson:"jenkins_url" json:"jenkins_url"`
	Name        string     `bson:"name" json:"name"`
	SdkType     string     `bson:"sdk_type" json:"sdk_type"`
	ProjectID   int32      `bson:"project_id" json:"project_id"`
	CreateTime  *time.Time `bson:"create_time" json:"create_time"`
	Shared      bool       `bson:"shared" json:"shared"`
}

// ChildSdk From admin_console
type ChildSdk struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SdkProjectID int32  `gorm:"column:sdk_project_id;not null" json:"sdk_project_id"` // SDK项目id
	ChildSdkName string `gorm:"column:child_sdk_name;not null" json:"child_sdk_name"` // 子sdk名称
	Jenkins      string `gorm:"column:jenkins;not null" json:"jenkins"`               // Jenkins
	Attribute    int32  `gorm:"column:attribute;not null" json:"attribute"`           // 属性（1：unity、2：iOS、3：Android、4：iOS配件）
	IsShare      int32  `gorm:"column:is_share;not null" json:"is_share"`             // 是否共享(0:否 1:是)
	CreatedAt    int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`   // 创建时间
	UpdatedAt    int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`   // 更新时间
	IsDeleted    bool   `gorm:"column:is_deleted;not null" json:"is_deleted"`         // 是否删除
}

func RunChildSdk() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("sdks")

	// 2、从mongo查询数据
	mSdks := make([]*MSdks, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mSdks)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	//fmt.Println(mSdks)

	// 3、将mongo数据装入切片
	childSdk := make([]*ChildSdk, 0)
	for i, sdk := range mSdks {
		fmt.Println("child_sdk: ", i)
		// share
		shared := 0
		if sdk.Shared == true {
			shared = 1
		}
		// 属性
		attribue := 0
		if sdk.SdkType == "unity" {
			attribue = 1
		} else if sdk.SdkType == "ios" {
			attribue = 2
		} else if sdk.SdkType == "android" {
			attribue = 3
		}
		cSdk := &ChildSdk{
			ID:           sdk.ID,
			SdkProjectID: sdk.ProjectID,
			ChildSdkName: sdk.Name,
			Jenkins:      sdk.JenkinsName,
			Attribute:    int32(attribue),
			IsShare:      int32(shared),
			CreatedAt:    sdk.CreateTime.Unix(),
			IsDeleted:    false,
		}
		childSdk = append(childSdk, cSdk)
	}

	// 4、将装有mongo数据的切片入库
	err = db2.MySQLClientAdmin.Table("child_sdk").CreateInBatches(childSdk, len(childSdk)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
