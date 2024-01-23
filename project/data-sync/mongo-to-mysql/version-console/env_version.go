package version_console

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/mao-gutils/constants"
	gj "github.com/mao888/mao-gutils/json"
	gutil "github.com/mao888/mao-gutils/strings"
	"go.mongodb.org/mongo-driver/bson"
)

// Env mapped from table version_console <env>
type Env struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`  // 主键
	AppID     int32  `gorm:"column:app_id;not null" json:"app_id"`               // 应用id
	Type      int32  `gorm:"column:type;not null" json:"type"`                   // 环境类型 0未知 1测试  3生产 4 自定义
	Name      string `gorm:"column:name;not null" json:"name"`                   // 环境名称
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 更新时间
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建时间
	IsDeleted int32  `gorm:"column:is_deleted;not null" json:"is_deleted"`       // 是否删除(0否1是)
}

func RunEnvAndVersion() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	collEnv := db.Collection("environments")
	collVersion := db.Collection("versions")

	// 2、从mongo查询数据
	// env
	mEnvironment := make([]*MEnvironment, 0)
	err := collEnv.Find(context.TODO(), bson.M{}).All(&mEnvironment)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	//fmt.Println(mEnvironment)

	// 3、将mongo数据装入切片
	//envs := make([]*Env, 0)
	countCha := 0
	countRu := 0
	for _, environment := range mEnvironment {

		arkEnvID := int(environment.EnvID)
		// 1、将mongo/environments -> mysql/env,并获得自增env.id
		// Type
		if environment.EnvID == 0 || environment.EnvID == 2 {
			continue
		}
		if environment.EnvID != 1 && environment.EnvID != 3 {
			environment.EnvID = 4
		}
		// IsDeleted
		isDeleted := 0
		if environment.DeleteTime != nil {
			isDeleted = 1
		}
		env := &Env{
			//ID:        0,
			AppID:     environment.AppID,
			Type:      environment.EnvID,
			Name:      environment.Name,
			UpdatedAt: environment.UpdateTime.Unix(),
			CreatedAt: environment.CreateTime.Unix(),
			IsDeleted: int32(isDeleted),
		}
		err = db2.MySQLClientVersion.Table("env").Create(env).Error
		if err != nil {
			fmt.Println("将mongo/environments -> mysql/env 错误", err)
			return
		}
		//fmt.Println("env.id = ", env.ID)

		// 2、根据 mysql/env app_id、type(原arkEnvID) 匹配 mongo/versions app_id、env_id 并mysql/env.id 并赋值 mongo/version.env_id 并入库mysql/version
		mVersion := make([]*MVersion, 0)
		err = collVersion.Find(context.TODO(), bson.M{"app_id": env.AppID, "env_id": arkEnvID}).All(&mVersion)
		if err != nil {
			fmt.Println("匹配 mongo/versions app_id、env_id错误：", err)
			return
		}
		countCha = countCha + len(mVersion)
		fmt.Printf("env: %d 下共有version数: %d \n", env.ID, countCha)

		// 3、如果mongo/versions parent_id 为空 入库mysql/version
		mapVersion := make(map[string]int32, 0)
		for _, version := range mVersion {
			fmt.Println("version1:", version)
			if version.ParentID != constants.EmptyString {
				continue
			}

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

			// IsGray 是否灰度  1是 0否
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
				EnvID:       env.ID,
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
			// mongo/versions parent_id 为空的 入库mysql/version
			err = db2.MySQLClientVersion.Table("version").Create(ver).Error
			if err != nil {
				fmt.Println("mongo/versions parent_id 为空的 入库mysql/version 错误：", err)
				return
			}
			fmt.Println("version1插入：", ver)
			countRu++
			mapVersion[version.ID] = ver.ID
		}
		//fmt.Printf("mapVersion := %v", mapVersion)
		fmt.Println("入库1:", countRu)

		// 4、mongo/versions parent_id 不为空 入库mysql/version
		for _, version := range mVersion {
			fmt.Println("version2:", version)
			if version.ParentID == constants.EmptyString {
				continue
			}
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
				ParentID:    mapVersion[version.ParentID],
				EnvID:       env.ID,
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
			// mongo/versions parent_id 不为空的 入库mysql/version
			err = db2.MySQLClientVersion.Table("version").Create(ver).Error
			if err != nil {
				fmt.Println("mongo/versions parent_id 不为空的 入库mysql/version 错误：", err)
				return
			}
			fmt.Println("version2插入:", env)
			countRu++
		}
		fmt.Println("入库2:", countRu)
	}
	fmt.Println("countCha:", countCha)
	fmt.Println("countRu:", countRu)
}
