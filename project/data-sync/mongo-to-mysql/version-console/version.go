package version_console

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	gj "github.com/mao888/mao-gutils/json"
	gutil "github.com/mao888/mao-gutils/strings"
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
	IsGray      int32  `gorm:"column:is_gray;not null" json:"is_gray"`             // 是否灰度 1是 0否
	GrayScale   int32  `gorm:"column:gray_scale;not null" json:"gray_scale"`       // 灰度范围 1 - 99
	Status      int32  `gorm:"column:status;not null" json:"status"`               // 版本状态 1未发布 2已发布 3已废弃
	PublishTime int64  `gorm:"column:publish_time;not null" json:"publish_time"`   // 发布时间
	Config      string `gorm:"column:config;not null" json:"config"`               // 版本配置，包括更新提示、全局配置、语言配置
	UpdatedAt   int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 更新时间
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建时间
	IsDeleted   int32  `gorm:"column:is_deleted;not null" json:"is_deleted"`       // 是否删除(0否1是)
}

// VersionConfig 版本配置，包括更新提示、全局配置、语言配置
type VersionConfig struct {
	Update *VersionConfigUpdate   `json:"update"` // 更新提示
	Global []*VersionConfigGlobal `json:"global"` // 全局配置	// GlobalConf
	Lang   []*VersionConfigLang   `json:"lang"`   // 语言配置
}
type VersionConfigUpdate struct {
	IsNotice    bool                       `json:"is_notice"`    // NoticeFlag
	LangType    int32                      `json:"lang_type"`    // MultiLngFlag 提示文案语言类型 1单一类型 2多种语言
	EnableClose bool                       `json:"enable_close"` // CloseFlag
	IsRestart   bool                       `json:"is_restart"`   // RestartFlag
	Text        []*VersionConfigUpdateText `json:"text"`
}
type VersionConfigUpdateText struct {
	Lang      string `json:"lang"`
	LangShort string `json:"lang_short"` // Lng
	IsDefault bool   `json:"is_default"` // DefaultLanguage
	Text      string `json:"text"`
}
type VersionConfigGlobal struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type VersionConfigLang struct {
	Lang      string                  `json:"lang"`
	LangShort string                  `json:"lang_short"` // Language
	IsDefault bool                    `json:"is_default"` // DefaultLng
	Args      []*VersionConfigLangArg `json:"args"`       /// ConfList
}
type VersionConfigLangArg struct {
	Key   string `json:"key"`
	Value string `json:"value"`
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

	//3、将mongo数据装入切片
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
		status := 0
		if version.Status == 2 {
			status = 1
		}
		if version.Status == 0 {
			status = 1
		} else if version.Status == 1 {
			status = 2
		} else if version.Status == 3 {
			status = 3
		}

		// IsGray 是否灰度 1是 0否
		isGray := 0
		if version.GrayFlag == true {
			isGray = 1
		}

		// IsDeleted 是否删除(0否1是)
		isDeleted := 0
		if version.DeleteTime != nil {
			isDeleted = 1
		}

		// PublishTime 发布时间
		var publishTime int64
		if version.PublishTime != nil {
			publishTime = version.PublishTime.Unix()
		} else if version.PublishTime == nil {
			publishTime = 0
		}

		// Config 版本配置，包括更新提示、全局配置、语言配置
		var config VersionConfig
		// 更新提示
		var versionConfigUpdate VersionConfigUpdate
		versionConfigUpdate.EnableClose = version.CloseFlag
		versionConfigUpdate.IsNotice = version.NoticeFlag
		versionConfigUpdate.IsRestart = version.RestartFlag
		if version.MultiLngFlag == true {
			versionConfigUpdate.LangType = 2
		} else if version.MultiLngFlag == false {
			versionConfigUpdate.LangType = 1
		}

		versionConfigUpdate.Text = make([]*VersionConfigUpdateText, 0)
		for _, text := range version.NoticeLngText {
			df := false
			if version.DefaultLanguage == text.Lng {
				df = true
			}
			t := &VersionConfigUpdateText{
				Lang:      "",
				LangShort: text.Lng,
				IsDefault: df,
				Text:      text.Text,
			}
			versionConfigUpdate.Text = append(versionConfigUpdate.Text, t)
		}
		config.Update = &versionConfigUpdate

		// 全局配置
		config.Global = make([]*VersionConfigGlobal, 0)
		for _, kv := range version.GlobalConf {
			versionConfigGlobal := &VersionConfigGlobal{
				Key:   kv.Key,
				Value: kv.Value,
			}
			config.Global = append(config.Global, versionConfigGlobal)
		}

		// 语言配置
		config.Lang = make([]*VersionConfigLang, 0)
		for _, conf := range version.LanguageConf {
			var lang VersionConfigLang

			lang.LangShort = conf.Language
			lang.IsDefault = conf.DefaultLng

			lang.Args = make([]*VersionConfigLangArg, 0)
			for _, kv := range conf.ConfList {
				k := &VersionConfigLangArg{
					Key:   kv.Key,
					Value: kv.Value,
				}
				lang.Args = append(lang.Args, k)
			}
			config.Lang = append(config.Lang, &lang)
		}

		configJson, err := gj.Object2JSONE(&config)
		if err != nil {
			fmt.Println(err)
			return
		}

		ver := &Version{
			//ID:          0,
			ParentID:    0,
			EnvID:       version.EnvID,
			Version:     version.VersionName,
			VersionNum:  gutil.VersionOrdinal(version.VersionName),
			Type:        int32(t),
			UpdateType:  version.UpdateType,
			IsGray:      int32(isGray),
			GrayScale:   version.GrayScale,
			Status:      int32(status),
			PublishTime: publishTime,
			Config:      configJson,
			UpdatedAt:   version.UpdateTime.Unix(),
			CreatedAt:   version.CreateTime.Unix(),
			IsDeleted:   int32(isDeleted),
		}
		versions = append(versions, ver)
	}

	// 4、将装有mongo数据的切片入库
	err = db2.MySQLClientVersion.Table("version").CreateInBatches(versions, len(versions)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}

	// 5、更新 ParentID
	//versionList := make([]*Version, 0)
	//err = db2.MySQLClientVersion.Table("version").Find(&versionList).Error
	//if err != nil {
	//	fmt.Println("mysql查询version错误：", err)
	//}
	//
	//for _, version := range versionList {
	//	// 根据 version.env_id 去env表查出app_id
	//	var env Env
	//	err := db2.MySQLClientVersion.Table("env").Where("id = ?", version.EnvID).First(&env).Error
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		fmt.Println("根据 version.env_id 去env表查出app_id,未查到")
	//	} else if err != nil {
	//		fmt.Println("根据 version.env_id 去env表查出app_id 错误：", err)
	//		return
	//	}
	//	// 根据app_id 和 env_id 去mongo查出 parent_id
	//	parentID := make([]*MVersion, 0)
	//	err = coll.Find(context.TODO(), bson.M{"app_id": env.AppID, "env_id": version.EnvID}).All(&parentID)
	//	if err != nil {
	//		fmt.Println("根据app_id 和 env_id 去mongo查出 parent_id 错误", err)
	//		return
	//	}
	//	// 如果 parent_id 为空，则无父版本id，跳过当前更新,默认为0
	//	if len(parentID) == 0 {
	//		continue
	//	}
	//
	//	// 根据ParentID从mongo查询父version信息
	//	parent := make([]*MVersion, 0)
	//	err = coll.Find(context.TODO(), bson.M{"_id": parentID[0].ParentID}).All(&parent)
	//	if err != nil {
	//		fmt.Println("根据parent_id Mongo查询错误", err)
	//		return
	//	}
	//	fmt.Println("parent:", len(parent))
	//
	//	// 根据 mongo父version中的 app_id 和 env_id 查出mysql父version
	//	var v Version
	//	err = db2.MySQLClientVersion.Table("version").
	//		Joins("inner join env on version.env_id = env.id").
	//		Where("version.env_id =? AND env.app_id = ?", parent[0].EnvID, parent[0].AppID).
	//		First(&v).Error
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		v.ParentID = 0
	//	} else if err != nil {
	//		fmt.Println("根据 mongo父version中的 app_id 和 env_id 查出mysql父version 错误：", err)
	//		return
	//	}
	//	// 更新mysql对应的 parent_id
	//	err = db2.MySQLClientVersion.Table("version").Where("id = ?", version.ID).
	//		UpdateColumn("parent_id", v.ID).Error
	//	if err != nil {
	//		fmt.Println("更新 ParentID 错误", err)
	//		return
	//	}
	//
	//}
}
