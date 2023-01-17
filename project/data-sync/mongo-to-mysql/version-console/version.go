package version_console

import (
	"context"
	"fmt"
	gutil "github.com/mao888/go-utils/strings"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
)

// Version mapped from table version_console <version>
type Version struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`  // 主键
	ParentID    int32  `gorm:"column:parent_id;not null" json:"parent_id"`         // 父版本id
	EnvID       int32  `gorm:"column:env_id;not null" json:"env_id"`               // 环境id
	Version     string `gorm:"column:version;not null" json:"version"`             // 版本号
	VersionNum  string `gorm:"column:version_num;not null" json:"version_num"`     // 字节序版本号
	Type        int32  `gorm:"column:type;not null" json:"type"`                   // 版本类型 1市场版本 2热更版本
	UpdateType  int32  `gorm:"column:update_type;not null" json:"update_type"`     // 更新类型 1强更 2非强更 3热更
	IsGray      int32  `gorm:"column:is_gray;not null" json:"is_gray"`             // 是否灰度 0未发布无灰度 1是 2否
	GrayScale   int32  `gorm:"column:gray_scale;not null" json:"gray_scale"`       // 灰度范围 1 - 99
	Status      int32  `gorm:"column:status;not null" json:"status"`               // 版本状态 1未发布 2已发布 3已废弃
	PublishTime int64  `gorm:"column:publish_time;not null" json:"publish_time"`   // 发布时间
	Config      string `gorm:"column:config;not null" json:"config"`               // 版本配置，包括更新提示、全局配置、语言配置
	UpdatedAt   int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 更新时间
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建时间
	IsDeleted   int32  `gorm:"column:is_deleted;not null" json:"is_deleted"`       // 是否删除(0否1是)
}

func RunVersion() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("versions")

	// 2、从mongo查询数据
	mVersion := make([]*MVersion, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mVersion)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mVersion)

	// 3、将mongo数据装入切片
	versions := make([]*Version, 0)
	for _, version := range mVersion {
		// type 版本类型 1市场版本 2热更版本
		t := 0
		if version.UpdateType == 1 || version.UpdateType == 2 {
			t = 1
		}
		if version.UpdateType == 3 {
			t = 2
		}

		// Status 版本状态 1未发布 2已发布 3已废弃
		if version.Status == 2 {
			continue
		}
		status := 0
		if version.Status == 0 {
			status = 1
		} else if version.Status == 1 {
			status = 2
		} else if version.Status == 3 {
			status = 3
		}

		// IsGray 是否灰度 0未发布无灰度 1是 2否
		isGray := 0
		if version.GrayFlag == true {
			isGray = 1
		} else if version.GrayFlag == false {
			isGray = 2
		}

		// IsDeleted 是否删除(0否1是)
		isDeleted := 0
		if version.DeleteTime != nil {
			isDeleted = 1
		}

		// ParentID 父版本id
		var parent MVersion
		err := coll.Find(context.TODO(), bson.M{"parent_id": version.ParentID}).All(&parent)
		if err != nil {
			fmt.Println("根据parent_id Mongo查询错误", err)
			return
		}
		var v Version
		err = db2.MySQLClientVersion.Table("version").
			Where("env_id = ?", parent.EnvID).
			Where("version = ?", parent.VersionName).
			Where("update_type = ?", parent.UpdateType).
			Where("status = ?", parent.Status).
			First(&v).Error
		if err != nil {
			fmt.Println("mysql查询id错误：", err)
		}

		ver := &Version{
			//ID:          0,
			ParentID:    v.ID,
			EnvID:       version.EnvID,
			Version:     version.VersionName,
			VersionNum:  gutil.VersionOrdinal(version.VersionName),
			Type:        int32(t),
			UpdateType:  version.UpdateType,
			IsGray:      int32(isGray),
			GrayScale:   version.GrayScale,
			Status:      int32(status),
			PublishTime: version.PublishTime.Unix(),
			Config:      "",
			UpdatedAt:   version.UpdateTime.Unix(),
			CreatedAt:   version.CreateTime.Unix(),
			IsDeleted:   int32(isDeleted),
		}
		versions = append(versions, ver)
	}

}
