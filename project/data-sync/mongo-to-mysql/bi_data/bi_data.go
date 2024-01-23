package bi_data

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	gerrors "github.com/mao888/mao-gerrors"
	gutil "github.com/mao888/mao-gutils/json"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// MThirdAdvConfigs From Mongo/rambler
type MThirdAdvConfigs struct {
	ID         string                   `bson:"_id" json:"_id"`
	Platform   string                   `bson:"platform" json:"platform"`
	AccList    []map[string]interface{} `bson:"acc_list" json:"acc_list"`
	OauthInfo  map[string]interface{}   `bson:"oauth_info,omitempty" json:"oauth_info,omitempty"`
	Comments   string                   `bson:"comments,omitempty" json:"comments,omitempty"`
	UpdateUser string                   `bson:"update_user,omitempty" json:"update_user,omitempty"`
	Type       string                   `bson:"type" json:"type"` // adv adn iap cloud
	DeleteTime *time.Time               `bson:"delete_time,omitempty" json:"-"`
	CreateTime *time.Time               `bson:"create_time" json:"create_time"`
	UpdateTime *time.Time               `bson:"update_time" json:"update_time"`
}

// BiDatum mapped from table admin_console <bi_data>
type BiDatum struct {
	ID            int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`    // 主键
	ThirdPlatform string `gorm:"column:third_platform;not null" json:"third_platform"` // 第三方平台
	DataDetails   string `gorm:"column:data_details;not null" json:"data_details"`     // 数据详情
	Type          string `gorm:"column:type;not null" json:"type"`                     // adv adn iap cloud
	CreatorID     int32  `gorm:"column:creator_id;not null" json:"creator_id"`         // 创建人id
	CreatorName   string `gorm:"column:creator_name;not null" json:"creator_name"`     // 创建人名称
	CreatedAt     int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`   // 创建时间
	UpdatedAt     int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`   // 更新时间
	IsDeleted     bool   `gorm:"column:is_deleted;not null" json:"is_deleted"`         // 是否删除
}

func RunBiData() {
	// 1、建立连接
	db := db2.MongoClient.Database("rambler")
	coll := db.Collection("thirdadvconfigs")

	// 2、从mongo查询数据
	mThirdAdvConfigs := make([]*MThirdAdvConfigs, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mThirdAdvConfigs)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mThirdAdvConfigs)

	// 3、将mongo数据装入切片
	biData := make([]*BiDatum, 0)
	for _, config := range mThirdAdvConfigs {

		// DataDetails
		accList, err := gutil.Object2JSONE(config.AccList)
		if err != nil {
			gerrors.Wrap(err, "Object2JSONE config.AccList err")
		}

		// IsDeleted
		isDeleted := false
		if config.DeleteTime != nil {
			isDeleted = true
		}
		data := &BiDatum{
			//ID:            0,
			ThirdPlatform: config.Platform,
			DataDetails:   accList,
			Type:          config.Type,
			//CreatorID:     0,
			CreatorName: config.UpdateUser,
			CreatedAt:   config.CreateTime.Unix(),
			UpdatedAt:   config.UpdateTime.Unix(),
			IsDeleted:   isDeleted,
		}
		biData = append(biData, data)
	}

	// 4、将装有mongo数据的切片入库
	err = db2.MySQLClientAdmin.Table("bi_data").CreateInBatches(biData, len(biData)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
