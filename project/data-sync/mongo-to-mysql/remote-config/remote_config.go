package remote_config

import (
	"context"
	"fmt"
	gutil "github.com/mao888/go-utils/json"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
)

// RemoteConfig mapped from table application_console <remote_config>
type RemoteConfig struct {
	ID         int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GameID     string `gorm:"column:game_id" json:"game_id"`         // 游戏ID
	Env        int32  `gorm:"column:env" json:"env"`                 // 环境: 1:master 2:test
	AppVersion string `gorm:"column:app_version" json:"app_version"` // 版本号
	GrayLevel  int32  `gorm:"column:gray_level" json:"gray_level"`   // 灰度
	Modules    string `gorm:"column:modules" json:"modules"`         // 功能模块
	Status     string `gorm:"column:status" json:"status"`           // 状态草稿、发布 stash、publish
	IsModified bool   `gorm:"column:is_modified" json:"is_modified"` // 是否有变更
	IsHidden   bool   `gorm:"column:is_hidden" json:"is_hidden"`     // 是否隐藏
	OriginID   int32  `gorm:"column:origin_id" json:"origin_id"`     // 原始id的拷贝
	Order      string `gorm:"column:order" json:"order"`             // 排序
	CreatorID  int32  `gorm:"column:creator_id" json:"creator_id"`   // 创建人ID
	CreatedAt  int64  `gorm:"column:created_at" json:"created_at"`   // 创建时间
	UpdatedAt  int64  `gorm:"column:updated_at" json:"updated_at"`   // 更新时间
	IsDeleted  bool   `gorm:"column:is_deleted" json:"is_deleted"`   // 是否删除
}

// MRemoteConfig From Mongo app_console <remote_config>
type MRemoteConfig struct {
	ID           string                   `bson:"_id" json:"id"`
	GameID       int32                    `bson:"game_id" json:"game_id"` // 游戏id
	GameIdCustom string                   `bson:"game_id_custom" json:"game_id_custom"`
	AppVersion   string                   `bson:"app_version" json:"app_version"` // 版本号
	Order        string                   `bson:"order" json:"order"`             // 排序
	GrayLevel    int32                    `bson:"gray_level" json:"gray_level"`   // 灰度
	Modules      []map[string]interface{} `bson:"modules" json:"modules"`         // 功能模块
	Status       string                   `bson:"status" json:"status"`           // 状态草稿、发布 stash、publish
	IsModified   bool                     `bson:"is_modified" json:"is_modified"` // 是否有变更
	IsHidden     bool                     `bson:"is_hidden" json:"is_hidden"`     // 是否隐藏
	//OriginID     string `bson:"origin_id" json:"origin_id"`     // 原始id的拷贝
	OperatorID int32 `bson:"operator_id" json:"operator_id"` // 创建人ID
	CreatedAt  int64 `bson:"created_at" json:"created_at"`   // 创建时间
	UpdatedAt  int64 `bson:"updated_at" json:"updated_at"`   // 更新时间
	IsDeleted  bool  `bson:"is_deleted" json:"is_deleted"`   // 是否删除
	Env        int32 `bson:"env" json:"env"`                 // 环境: 1:master 2:test
}

func RunRemoteConfig() {
	// 1、建立连接
	db := db2.MongoClient.Database("app_console")
	coll := db.Collection("remote_config")

	// 2、从mongo查询数据
	mRemoteConfig := make([]*MRemoteConfig, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mRemoteConfig)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mRemoteConfig)

	// 3、将mongo数据装入切片
	remoteConfig := make([]*RemoteConfig, 0)
	for _, config := range mRemoteConfig {
		if config.Status == "stash" {
			continue
		}
		modules, err := gutil.Object2JSONE(&config.Modules)
		if err != nil {
			return
		}
		remote := &RemoteConfig{
			//ID:         0,
			GameID:     config.GameIdCustom,
			Env:        config.Env,
			AppVersion: config.AppVersion,
			GrayLevel:  config.GrayLevel,
			Modules:    modules,
			Status:     config.Status,
			IsModified: config.IsModified,
			IsHidden:   config.IsHidden,
			OriginID:   0,
			Order:      config.Order,
			CreatorID:  config.OperatorID,
			CreatedAt:  config.CreatedAt,
			UpdatedAt:  config.UpdatedAt,
			IsDeleted:  config.IsDeleted,
		}
		remoteConfig = append(remoteConfig, remote)
	}

	// 4、将装有mongo数据的切片入库
	err = db2.MySQLClient.Table("remote_config").CreateInBatches(remoteConfig, len(remoteConfig)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
