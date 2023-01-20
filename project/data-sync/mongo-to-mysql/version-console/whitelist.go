package version_console

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
)

// Whitelist mapped from table version_console <whitelist>
type Whitelist struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`  // 主键
	AppID     int32  `gorm:"column:app_id;not null" json:"app_id"`               // 应用id
	DeviceID  string `gorm:"column:device_id;not null" json:"device_id"`         // 设备id
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 更新时间
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建时间
	IsDeleted int32  `gorm:"column:is_deleted;not null" json:"is_deleted"`       // 是否删除(0否1是)
}

func RunWhitelist() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("whitelists")

	// 2、从mongo查询数据
	mWhiteList := make([]*MWhiteList, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mWhiteList)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mWhiteList)

	// 3、将mongo数据装入切片
	whiteList := make([]*Whitelist, 0)
	for _, list := range mWhiteList {
		deleted := 0
		if list.DeleteTime != nil {
			deleted = 1
		}
		white := &Whitelist{
			//ID:        0,
			AppID:    list.AppID,
			DeviceID: list.DevCode,
			//UpdatedAt: 0,
			CreatedAt: list.CreateTime.Unix(),
			IsDeleted: int32(deleted),
		}
		whiteList = append(whiteList, white)
	}
	// 4、将装有mongo数据的切片入库
	err = db2.MySQLClientVersion.Table("whitelist").CreateInBatches(whiteList, len(whiteList)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
