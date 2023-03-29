package sdk_release

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// MProjects From Mongo
type MProjects struct {
	ID         int32      `bson:"_id" json:"_id"`
	Show       bool       `bson:"show" json:"show"`
	Name       string     `bson:"name" json:"name"`
	NameEn     string     `bson:"name_en" json:"name_en"`
	CreateTime *time.Time `bson:"create_time" json:"create_time"`
	UpdateTime *time.Time `bson:"update_time" json:"update_time"`
}

// SdkProject From admin_console
type SdkProject struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Sdk         string `gorm:"column:sdk;not null" json:"sdk"`                     // SDK项目（中文名称）
	EnglishName string `gorm:"column:english_name;not null" json:"english_name"`   // 英文名称
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建时间
	UpdatedAt   int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 更新时间
	IsDeleted   bool   `gorm:"column:is_deleted;not null" json:"is_deleted"`       // 是否删除
	IsGia       bool   `gorm:"column:is_gia;not null" json:"is_gia"`               // 对gia显示(0:关闭 1:开启)
}

func RunSdkProject() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("projects")

	// 2、从mongo查询数据
	mProjects := make([]*MProjects, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mProjects)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	//fmt.Println(mProjects)

	// 3、将mongo数据装入切片
	sdkProject := make([]*SdkProject, 0)
	for i, project := range mProjects {
		fmt.Println("sdk_project: ", i)
		sdk := &SdkProject{
			ID:          project.ID,
			Sdk:         project.Name,
			EnglishName: project.NameEn,
			CreatedAt:   project.CreateTime.Unix(),
			UpdatedAt:   project.UpdateTime.Unix(),
			IsDeleted:   false,
			IsGia:       project.Show,
		}
		sdkProject = append(sdkProject, sdk)
	}

	// 4、将装有mongo数据的切片入库
	err = db2.MySQLClientAdmin.Table("sdk_project").CreateInBatches(sdkProject, len(sdkProject)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
