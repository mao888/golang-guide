package mysql_to_mysql

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// MGameConfigs From Mongo
type MGameConfigs struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	CompanyID     int32              `bson:"company_id" json:"company_id"`
	GameID        int32              `bson:"game_id" json:"game_id"`                         //自定义游戏Id，唯一
	CfgType       int32              `bson:"cfg_type" json:"cfg_type,omitempty"`             // 配置类型 1: 邮件管理配置。2：活动管理配置
	TargetService int32              `bson:"target_service" json:"target_service,omitempty"` //目标服务器 1：测试服。3：正式服
	Url           string             `bson:"url" json:"url"`                                 //服务器url
	Operator      int32              `bson:"operator" json:"operator,omitempty"`             // 操作人
	CreateTime    *time.Time         `bson:"create_time" json:"create_time"`
	UpdateTime    *time.Time         `bson:"update_time" json:"update_time"`
}

func RunMGameConfigs() {
	// 1、建立Mongo连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("gameconfigs")
	collGame := db.Collection("games")

	// 2、从mongo查询数据
	var mGameConfigs []MGameConfigs
	err := coll.Find(context.TODO(), bson.M{}).All(&mGameConfigs)
	if err != nil {
		fmt.Println("Mongo查询 gameconfigs 错误：", err)
	}

	var mGame []MGame
	err = collGame.Find(context.TODO(), bson.M{}).All(&mGame)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
	}

	// 将从mongo中查出的games.id(int)作为key, games.game_id(string)作为value,存入map
	idMap := map[int32]string{}
	for _, game := range mGame {
		idMap[game.ID] = game.GameID
	}

	// 3、将mongo数据装入切片
	gmConfig := make([]*GmConfig, 0)
	for _, config := range mGameConfigs {
		// env 1 正式环境 2 测试环境
		if config.TargetService == 1 {
			config.TargetService = 2
		}
		if config.TargetService == 3 {
			config.TargetService = 1
		}
		now := time.Now()
		if config.CreateTime == nil {
			config.CreateTime = &now
		}
		if config.UpdateTime == nil {
			config.UpdateTime = &now
		}
		gm := &GmConfig{
			// ID:        config.ID,
			GameID:    idMap[config.GameID],
			Env:       config.TargetService,
			Type:      config.CfgType,
			URL:       config.Url,
			CreatorID: config.Operator,
			CreatedAt: config.CreateTime.Unix(),
			UpdatedAt: config.UpdateTime.Unix(),
		}
		gmConfig = append(gmConfig, gm)
	}
	fmt.Println(gmConfig)
	err = db2.MySQLClient.Table("gm_config").CreateInBatches(gmConfig, len(gmConfig)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
