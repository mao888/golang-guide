package version_console

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
)

// Language mapped from table version_console <language>
type Language struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`  // 主键
	AppID     int32  `gorm:"column:app_id;not null" json:"app_id"`               // 应用id
	Lang      string `gorm:"column:lang;not null" json:"lang"`                   // 语言名
	LangShort string `gorm:"column:lang_short;not null" json:"lang_short"`       // 语言名缩写
	IsDefault int32  `gorm:"column:is_default;not null" json:"is_default"`       // 是否默认语言 0否1是
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 更新时间
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建时间
	IsDeleted int32  `gorm:"column:is_deleted;not null" json:"is_deleted"`       // 是否删除(0否1是)
}

func RunLanguage() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("gamelanguageconfs")

	// 2、从mongo查询数据
	mGameLanguageConf := make([]*MGameLanguageConf, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mGameLanguageConf)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mGameLanguageConf)

	// 3、将mongo数据装入切片
	language := make([]*Language, 0)
	for _, conf := range mGameLanguageConf {
		//fmt.Println(conf.AppID)
		isDefault := 0
		if conf.DefaultLng == true {
			isDefault = 1
		}
		lan := &Language{
			//ID:        0,
			AppID:     conf.AppID,
			Lang:      conf.NameEn,
			LangShort: conf.NameShort,
			IsDefault: int32(isDefault),
			UpdatedAt: conf.UpdateTime.Unix(),
			CreatedAt: conf.CreateTime.Unix(),
		}
		language = append(language, lan)
	}

	// 4、将装有mongo数据的切片入库
	err = db2.MySQLClientVersion.Table("language").CreateInBatches(language, len(language)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
